/*
*	Copyright (C) 2022  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package client

import (
	"context"

	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sResourceClient struct {
	client *kubernetes.Clientset
}

type Service struct {
	URL         string `json:"url" yaml:"url"`
	Description string `json:"description" yaml:"description"`
	SwaggerURL  string `json:"swagger" yaml:"swagger"`
	Uptime      string `json:"uptime" yaml:"uptime"`
}

func (k8s *K8sResourceClient) initializeK8sClient() (*kubernetes.Clientset, error) {
	if config, e := rest.InClusterConfig(); e == nil {
		return kubernetes.NewForConfig(config)
	} else {
		return nil, e
	}
}

func (k8s *K8sResourceClient) getClusterIngresses() (*v1.IngressList, error) {
	return k8s.client.NetworkingV1().Ingresses("").List(context.Background(), metav1.ListOptions{
		Limit: 256,
	})
}

func (k8s *K8sResourceClient) CollectServices() ([]*Service, error) {
	if ilist, e := k8s.getClusterIngresses(); e != nil || ilist == nil {
		return nil, e
	} else {
		// var services []*Service

		// for _, ingress := range ilist.Items {

		// }

	}
}
