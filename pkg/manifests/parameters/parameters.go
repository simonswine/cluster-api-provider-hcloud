package parameters

import (
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ManifestParameters struct {
	HcloudToken         *string
	RobotUserName       *string
	RobotPassword       *string
	HcloudNetwork       *intstr.IntOrString
	KubeAPIServerIPv4   *string
	KubeAPIServerDomain *string
	Port                *string
	CAcrt               *string
	CAkey               *string
}

func (m *ManifestParameters) ExtVar() map[string]string {
	extVar := make(map[string]string)

	if key, val := "kube-apiserver-ip", m.KubeAPIServerIPv4; val != nil {
		extVar[key] = *val
	} else {
		extVar[key] = ""
	}

	if key, val := "kube-apiserver-domain", m.KubeAPIServerDomain; val != nil {
		extVar[key] = *val
	} else {
		extVar[key] = ""
	}

	if key, val := "port", m.Port; val != nil {
		extVar[key] = *val
	} else {
		extVar[key] = "6443"
	}

	if key, val := "hcloud-token", m.HcloudToken; val != nil {
		extVar[key] = *val
	}

	if key, val := "robot-username", m.RobotUserName; val != nil {
		extVar[key] = *val
	}

	if key, val := "robot-password", m.RobotPassword; val != nil {
		extVar[key] = *val
	}

	if key, val := "ca-crt", m.CAcrt; val != nil {
		extVar[key] = *val
	}

	if key, val := "ca-key", m.CAkey; val != nil {
		extVar[key] = *val
	}

	if key, val := "hcloud-network", m.HcloudNetwork; val != nil {
		extVar[key] = val.String()
	} else {
		extVar[key] = ""
	}

	return extVar
}
