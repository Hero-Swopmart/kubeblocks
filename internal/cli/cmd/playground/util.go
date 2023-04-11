/*
Copyright ApeCloud, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package playground

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	cp "github.com/apecloud/kubeblocks/internal/cli/cloudprovider"
	"github.com/apecloud/kubeblocks/internal/cli/printer"
	"github.com/apecloud/kubeblocks/internal/cli/util"
	"github.com/apecloud/kubeblocks/version"
)

func playgroundDir() (string, error) {
	cliPath, err := util.GetCliHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cliPath, "playground"), nil
}

// cloudProviderRepoDir cloud provider repo directory
func cloudProviderRepoDir() (string, error) {
	dir, err := playgroundDir()
	if err != nil {
		return "", err
	}
	major := strings.Split(version.Version, "-")[0]
	cpDir := cp.GitRepoName
	if major != "" {
		cpDir = fmt.Sprintf("%s-%s", cp.GitRepoName, major)
	}
	return filepath.Join(dir, cpDir), err
}

func initPlaygroundDir() (string, error) {
	dir, err := playgroundDir()
	if err != nil {
		return "", err
	}

	if _, err = os.Stat(dir); err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0750)
	}
	return dir, err
}

// writeClusterInfo writes the cluster info to a state file
func writeClusterInfo(path string, info *cp.K8sClusterInfo) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0640)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = json.NewEncoder(f).Encode(info); err != nil {
		// if we fail to write the state file, we should remove it
		if e := os.Remove(path); e != nil {
			return errors.Wrap(err, e.Error())
		}
		return err
	}
	return nil
}

func removeStateFile(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return os.Remove(path)
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// readClusterInfoFromFile reads the kubernetes cluster info from a state file
func readClusterInfoFromFile(path string) (*cp.K8sClusterInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	var info cp.K8sClusterInfo
	if err = json.NewDecoder(f).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}

func writeAndUseKubeConfig(kubeConfig string, kubeConfigPath string, out io.Writer) error {
	spinner := printer.Spinner(out, fmt.Sprintf("%-50s", "Write kubeconfig to "+kubeConfigPath))
	defer spinner(false)
	if err := kubeConfigWrite(kubeConfig, kubeConfigPath, writeKubeConfigOptions{
		UpdateExisting:       true,
		UpdateCurrentContext: true,
		OverwriteExisting:    true}); err != nil {
		return err
	}

	// use the new kubeconfig file
	if err := util.SetKubeConfig(kubeConfigPath); err != nil {
		return err
	}

	spinner(true)
	return nil
}

// getKubeClient returns a kubernetes dynamic client and check if the cluster is reachable
func getKubeClient() (kubernetes.Interface, dynamic.Interface, error) {
	f := util.NewFactory()
	client, err := f.KubernetesClientSet()
	errMsg := kubeClusterUnreachableErr.Error()
	if err == genericclioptions.ErrEmptyConfig {
		return nil, nil, kubeClusterUnreachableErr
	}
	if err != nil {
		return nil, nil, errors.Wrap(err, errMsg)
	}

	if _, err = client.ServerVersion(); err != nil {
		return nil, nil, errors.Wrap(err, errMsg)
	}

	dynamic, err := f.DynamicClient()
	if err != nil {
		return nil, nil, errors.Wrap(err, errMsg)
	}
	return client, dynamic, nil
}
