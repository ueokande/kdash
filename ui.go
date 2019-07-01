package main

import (
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

const maxPaneWidth = 40
const paneWidthRatio = 0.2

type PaneType int

const (
	ContextPane PaneType = iota
	PodsPane
	ServicesPane
	DeploymentsPane
	InfoPane
)

type Style struct {
	SelectedRowStyle   termui.Style
	UnselectedRowStyle termui.Style
}

type UI struct {
	pods        []corev1.Pod
	services    []corev1.Service
	deployments []appsv1.Deployment

	contextWidget     *widgets.Paragraph
	podsWidget        *widgets.List
	servicesWidget    *widgets.List
	deploymentsWidget *widgets.List
	infoWidget        *widgets.Paragraph

	style Style

	activePane PaneType
}

func NewUI() *UI {
	style := Style{
		SelectedRowStyle: termui.Style{
			Bg: termui.ColorWhite,
			Fg: termui.ColorBlack,
		},
		UnselectedRowStyle: termui.Style{
			Bg: termui.ColorClear,
			Fg: termui.ColorClear,
		},
	}

	contextWidget := widgets.NewParagraph()
	contextWidget.Title = "Cluster/namespace"
	contextWidget.TextStyle = termui.NewStyle(termui.ColorBlue)

	podsWidget := widgets.NewList()
	podsWidget.Title = "Pods"
	podsWidget.TextStyle = termui.NewStyle(termui.ColorYellow)
	podsWidget.WrapText = false
	podsWidget.TextStyle = style.UnselectedRowStyle
	podsWidget.SelectedRowStyle = style.SelectedRowStyle

	servicesWidget := widgets.NewList()
	servicesWidget.Title = "Services"
	servicesWidget.TextStyle = termui.NewStyle(termui.ColorYellow)
	servicesWidget.WrapText = false
	servicesWidget.TextStyle = style.UnselectedRowStyle
	servicesWidget.SelectedRowStyle = style.SelectedRowStyle

	deploymentsWidget := widgets.NewList()
	deploymentsWidget.Title = "Deployments"
	deploymentsWidget.TextStyle = termui.NewStyle(termui.ColorYellow)
	deploymentsWidget.WrapText = false
	deploymentsWidget.TextStyle = style.UnselectedRowStyle
	deploymentsWidget.SelectedRowStyle = style.SelectedRowStyle

	infoWidget := widgets.NewParagraph()
	infoWidget.Title = "Info"

	ui := &UI{
		contextWidget:     contextWidget,
		podsWidget:        podsWidget,
		servicesWidget:    servicesWidget,
		deploymentsWidget: deploymentsWidget,
		infoWidget:        infoWidget,
		activePane:        PodsPane,

		style: style,
	}
	ui.selectActivePane()
	ui.Resize()
	return ui
}

func (u *UI) SelectNextPane() {
	u.activePane = func() PaneType {
		switch u.activePane {
		case ContextPane:
			return PodsPane
		case PodsPane:
			return ServicesPane
		case ServicesPane:
			return DeploymentsPane
		case DeploymentsPane:
			return InfoPane
		case InfoPane:
			return PodsPane
		}
		return PodsPane
	}()
	u.selectActivePane()
}

func (u *UI) SelectPrevPane() {
	u.activePane = func() PaneType {
		switch u.activePane {
		case InfoPane:
			return DeploymentsPane
		case DeploymentsPane:
			return ServicesPane
		case ServicesPane:
			return PodsPane
		case PodsPane:
			return ContextPane
		case ContextPane:
			return InfoPane
		}
		return PodsPane
	}()
	u.selectActivePane()
}

func (u *UI) selectActivePane() {
	u.contextWidget.BorderStyle.Fg = termui.ColorBlack
	u.podsWidget.BorderStyle.Fg = termui.ColorBlack
	u.podsWidget.SelectedRowStyle = u.style.UnselectedRowStyle
	u.servicesWidget.BorderStyle.Fg = termui.ColorBlack
	u.servicesWidget.SelectedRowStyle = u.style.UnselectedRowStyle
	u.deploymentsWidget.BorderStyle.Fg = termui.ColorBlack
	u.deploymentsWidget.SelectedRowStyle = u.style.UnselectedRowStyle
	u.infoWidget.BorderStyle.Fg = termui.ColorBlack

	switch u.activePane {
	case ContextPane:
		u.contextWidget.BorderStyle.Fg = termui.ColorCyan
	case PodsPane:
		u.podsWidget.BorderStyle.Fg = termui.ColorCyan
		u.podsWidget.SelectedRowStyle = u.style.SelectedRowStyle
	case ServicesPane:
		u.servicesWidget.BorderStyle.Fg = termui.ColorCyan
		u.servicesWidget.SelectedRowStyle = u.style.SelectedRowStyle
	case DeploymentsPane:
		u.deploymentsWidget.BorderStyle.Fg = termui.ColorCyan
		u.deploymentsWidget.SelectedRowStyle = u.style.SelectedRowStyle
	case InfoPane:
		u.infoWidget.BorderStyle.Fg = termui.ColorCyan
	}
}

func (u *UI) ScrollDown() {
	switch u.activePane {
	case PodsPane:
		u.podsWidget.ScrollDown()
	case ServicesPane:
		u.servicesWidget.ScrollDown()
	case DeploymentsPane:
		u.deploymentsWidget.ScrollDown()
	}
}

func (u *UI) ScrollUp() {
	switch u.activePane {
	case PodsPane:
		u.podsWidget.ScrollUp()
	case ServicesPane:
		u.servicesWidget.ScrollUp()
	case DeploymentsPane:
		u.deploymentsWidget.ScrollUp()
	}
}

func (u *UI) ScrollHalfPageDown() {
	switch u.activePane {
	case PodsPane:
		u.podsWidget.ScrollHalfPageDown()
	case ServicesPane:
		u.servicesWidget.ScrollHalfPageDown()
	case DeploymentsPane:
		u.deploymentsWidget.ScrollHalfPageDown()
	}
}

func (u *UI) ScrollHalfPageUp() {
	switch u.activePane {
	case PodsPane:
		u.podsWidget.ScrollHalfPageUp()
	case ServicesPane:
		u.servicesWidget.ScrollHalfPageUp()
	case DeploymentsPane:
		u.deploymentsWidget.ScrollHalfPageUp()
	}
}

func (u *UI) ScrollPageDown() {
	switch u.activePane {
	case PodsPane:
		u.podsWidget.ScrollPageDown()
	case ServicesPane:
		u.servicesWidget.ScrollPageDown()
	case DeploymentsPane:
		u.deploymentsWidget.ScrollPageDown()
	}
}

func (u *UI) ScrollPageUp() {
	switch u.activePane {
	case PodsPane:
		u.podsWidget.ScrollPageUp()
	case ServicesPane:
		u.servicesWidget.ScrollPageUp()
	case DeploymentsPane:
		u.deploymentsWidget.ScrollPageUp()
	}
}

func (u *UI) ScrollTop() {
	switch u.activePane {
	case PodsPane:
		u.podsWidget.ScrollTop()
	case ServicesPane:
		u.servicesWidget.ScrollTop()
	case DeploymentsPane:
		u.deploymentsWidget.ScrollTop()
	}
}

func (u *UI) ScrollBottom() {
	switch u.activePane {
	case PodsPane:
		u.podsWidget.ScrollBottom()
	case ServicesPane:
		u.servicesWidget.ScrollBottom()
	case DeploymentsPane:
		u.deploymentsWidget.ScrollBottom()
	}
}

func (u *UI) SetContext(cluster, namespace string) {
	u.contextWidget.Text = cluster + "/" + namespace
}

func (u *UI) SetPods(pods []corev1.Pod) {
	u.pods = pods

	rows := make([]string, len(u.pods))
	for i, p := range u.pods {
		rows[i] = p.ObjectMeta.Name
	}
	u.podsWidget.Rows = rows
}

func (u *UI) SetServices(services []corev1.Service) {
	u.services = services

	rows := make([]string, len(u.services))
	for i, p := range u.services {
		rows[i] = p.ObjectMeta.Name
	}
	u.servicesWidget.Rows = rows
}

func (u *UI) SetDeployments(deployments []appsv1.Deployment) {
	u.deployments = deployments

	rows := make([]string, len(u.deployments))
	for i, p := range u.deployments {
		rows[i] = p.ObjectMeta.Name
	}
	u.deploymentsWidget.Rows = rows
}

func (u *UI) Resize() {
	termw, termh := termui.TerminalDimensions()
	paneWidth := int(float64(termw) * paneWidthRatio)
	if paneWidth > maxPaneWidth {
		paneWidth = maxPaneWidth
	}

	u.contextWidget.SetRect(0, 0, paneWidth, 3)
	u.podsWidget.SetRect(0, 3, paneWidth, 20)
	u.servicesWidget.SetRect(0, 20, paneWidth, 30)
	u.deploymentsWidget.SetRect(0, 30, paneWidth, termh)
	u.infoWidget.SetRect(paneWidth, 0, termw, termh)

	u.Render()
}

func (u *UI) Render() {
	u.setInfo()

	termui.Render(u.contextWidget,
		u.podsWidget,
		u.servicesWidget,
		u.deploymentsWidget,
		u.infoWidget)
}

func (u *UI) setInfo() {
	u.infoWidget.Text = `
Name:  nginx-deployment-1006230814-6winp
Namespace: default
Node:  kubernetes-node-wul5/10.240.0.9
Start Time: Thu, 24 Mar 2016 01:39:49 +0000
Labels:  app=nginx,pod-template-hash=1006230814
Annotations:    kubernetes.io/created-by={"kind":"SerializedReference","apiVersion":"v1","reference":{"kind":"ReplicaSet","namespace":"default","name":"nginx-deployment-1956810328","uid":"14e607e7-8ba1-11e7-b5cb-fa16" ...
Status:  Running
IP:  10.244.0.6
Controllers: ReplicaSet/nginx-deployment-1006230814
Containers:
  nginx:
    Container ID: docker://90315cc9f513c724e9957a4788d3e625a078de84750f244a40f97ae355eb1149
    Image:  nginx
    Image ID:  docker://6f62f48c4e55d700cf3eb1b5e33fa051802986b77b874cc351cce539e5163707
    Port:  80/TCP
    QoS Tier:
      cpu: Guaranteed
      memory: Guaranteed
    Limits:
      cpu: 500m
      memory: 128Mi
    Requests:
      memory:  128Mi
      cpu:  500m
    State:  Running
      Started:  Thu, 24 Mar 2016 01:39:51 +0000
    Ready:  True
    Restart Count: 0
    Environment:        <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-5kdvl (ro)
Conditions:
  Type          Status
  Initialized   True
  Ready         True
  PodScheduled  True
Volumes:
  default-token-4bcbi:
    Type: Secret (a volume populated by a Secret)
    SecretName: default-token-4bcbi
    Optional:   false
QoS Class:      Guaranteed
Node-Selectors: <none>
Tolerations:    <none>
Events:
  FirstSeen LastSeen Count From     SubobjectPath  Type  Reason  Message
  --------- -------- ----- ----     -------------  -------- ------  -------
  54s  54s  1 {default-scheduler }      Normal  Scheduled Successfully assigned nginx-deployment-1006230814-6winp to kubernetes-node-wul5
  54s  54s  1 {kubelet kubernetes-node-wul5} spec.containers{nginx} Normal  Pulling  pulling image "nginx"
  53s  53s  1 {kubelet kubernetes-node-wul5} spec.containers{nginx} Normal  Pulled  Successfully pulled image "nginx"
  53s  53s  1 {kubelet kubernetes-node-wul5} spec.containers{nginx} Normal  Created  Created container with docker id 90315cc9f513
  53s  53s  1 {kubelet kubernetes-node-wul5} spec.containers{nginx} Normal  Started  Started container with docker id 90315cc9f513
`
}
