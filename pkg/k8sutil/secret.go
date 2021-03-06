/*
Copyright 2019 Banzai Cloud.

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

package k8sutil

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/goph/emperror"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetSercretByName(name, namespace string, client client.Client, log logr.Logger) (*corev1.Secret, error) {
	key := types.NamespacedName{
		Name:      name,
		Namespace: namespace,
	}
	log.Info("secret getter", "name", name, "namespace", namespace)
	var secret corev1.Secret
	if err := client.Get(context.TODO(), key, &secret); err != nil {
		return nil, emperror.Wrap(err, "cannot get secret by name")
	}

	return &secret, nil
}
