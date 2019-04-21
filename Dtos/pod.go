package Dtos

// Pod json structure to read from kubernetes api server to detect what's pods' status are. after connecting to
// kubernetes api server this structure fulfilled and traversed against number of running pods
// to satisfy the number of active pods to exit unless timeout timer reaches if it was set

type Pod struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		Metadata struct {
			CreationTimestamp string `json:"creationTimestamp"`
			GenerateName      string `json:"generateName"`
			Labels            struct {
				App               string `json:"app"`
				Pod_template_hash string `json:"pod-template-hash"`
			} `json:"labels"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			OwnerReferences []struct {
				APIVersion         string `json:"apiVersion"`
				BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
				Controller         bool   `json:"controller"`
				Kind               string `json:"kind"`
				Name               string `json:"name"`
				UID                string `json:"uid"`
			} `json:"ownerReferences"`
			ResourceVersion string `json:"resourceVersion"`
			SelfLink        string `json:"selfLink"`
			UID             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			Containers []struct {
				Args                     []string `json:"args"`
				Command                  []string `json:"command"`
				Image                    string   `json:"image"`
				ImagePullPolicy          string   `json:"imagePullPolicy"`
				Name                     string   `json:"name"`
				Resources                struct{} `json:"resources"`
				TerminationMessagePath   string   `json:"terminationMessagePath"`
				TerminationMessagePolicy string   `json:"terminationMessagePolicy"`
				VolumeMounts             []struct {
					MountPath string `json:"mountPath"`
					Name      string `json:"name"`
					ReadOnly  bool   `json:"readOnly"`
				} `json:"volumeMounts"`
			} `json:"containers"`
			DNSPolicy          string `json:"dnsPolicy"`
			EnableServiceLinks bool   `json:"enableServiceLinks"`
			NodeName           string `json:"nodeName"`
			NodeSelector       struct {
				Kubernetes_io_hostname string `json:"kubernetes.io/hostname"`
			} `json:"nodeSelector"`
			Priority                      int      `json:"priority"`
			RestartPolicy                 string   `json:"restartPolicy"`
			SchedulerName                 string   `json:"schedulerName"`
			SecurityContext               struct{} `json:"securityContext"`
			ServiceAccount                string   `json:"serviceAccount"`
			ServiceAccountName            string   `json:"serviceAccountName"`
			TerminationGracePeriodSeconds int      `json:"terminationGracePeriodSeconds"`
			Tolerations                   []struct {
				Effect            string `json:"effect"`
				Key               string `json:"key"`
				Operator          string `json:"operator"`
				TolerationSeconds int    `json:"tolerationSeconds"`
			} `json:"tolerations"`
			Volumes []struct {
				Name   string `json:"name"`
				Secret struct {
					DefaultMode int    `json:"defaultMode"`
					SecretName  string `json:"secretName"`
				} `json:"secret"`
			} `json:"volumes"`
		} `json:"spec"`
		Status struct {
			Conditions []struct {
				LastProbeTime      interface{} `json:"lastProbeTime"`
				LastTransitionTime string      `json:"lastTransitionTime"`
				Status             string      `json:"status"`
				Type               string      `json:"type"`
			} `json:"conditions"`
			ContainerStatuses []struct {
				ContainerID string `json:"containerID"`
				Image       string `json:"image"`
				ImageID     string `json:"imageID"`
				LastState   struct {
					Terminated struct {
						ContainerID string `json:"containerID"`
						ExitCode    int    `json:"exitCode"`
						FinishedAt  string `json:"finishedAt"`
						Reason      string `json:"reason"`
						StartedAt   string `json:"startedAt"`
					} `json:"terminated"`
				} `json:"lastState"`
				Name         string `json:"name"`
				Ready        bool   `json:"ready"`
				RestartCount int    `json:"restartCount"`
				State        struct {
					Running struct {
						StartedAt string `json:"startedAt"`
					} `json:"running"`
				} `json:"state"`
			} `json:"containerStatuses"`
			HostIP    string `json:"hostIP"`
			Phase     string `json:"phase"`
			PodIP     string `json:"podIP"`
			QosClass  string `json:"qosClass"`
			StartTime string `json:"startTime"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
	} `json:"metadata"`
}
