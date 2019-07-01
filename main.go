package main

import (
	"fmt"
	"os"

	termui "github.com/gizak/termui/v3"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func run() error {
	if err := termui.Init(); err != nil {
		return err
	}
	defer termui.Close()

	ui := NewUI()
	ui.SetContext("my-service.cluster.local", "default")
	ui.SetPods([]corev1.Pod{
		{ObjectMeta: metav1.ObjectMeta{Name: "front-app-abcd1234-abcdef"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "front-app-abcd1234-019bbi"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "front-app-abcd1234-xb4k10"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "front-app-abcd1234-zb8kk3"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "front-app-abcd1234-487bk2"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "auth-app-u90211bs-9b811a"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "auth-app-u90211bs-76bny3"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "auth-app-u90211bs-mvkq03"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "async-queue-9nb019f8-90g935"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "async-queue-9nb019f8-bjdk54"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "async-queue-9nb019f8-eetjb9"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "job-scheduler-b913mgl3-f92y4j"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "job-scheduler-b913mgl3-019bl1"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "job-scheduler-b913mgl3-912333"}},
	})
	ui.SetServices([]corev1.Service{
		{ObjectMeta: metav1.ObjectMeta{Name: "front-app"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "auth-app"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "async-queue"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "job-scheduler"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "front-rds"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "job-rds"}},
	})
	ui.SetDeployments([]appsv1.Deployment{
		{ObjectMeta: metav1.ObjectMeta{Name: "front-app"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "auth-app"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "async-queue"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "job-scheduler"}},
	})
	ui.Resize()
	ui.Render()

	uiEvents := termui.PollEvents()
	prevKey := ""
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-n>":
			ui.SelectNextPane()
		case "<C-p>":
			ui.SelectPrevPane()
		case "j":
			ui.ScrollDown()
		case "k":
			ui.ScrollUp()
		case "<C-d>":
			ui.ScrollHalfPageDown()
		case "<C-u>":
			ui.ScrollHalfPageUp()
		case "<C-f>":
			ui.ScrollPageDown()
		case "<C-b>":
			ui.ScrollPageUp()
		case "g":
			if prevKey == "g" {
				ui.ScrollTop()
			}
		case "G":
			ui.ScrollBottom()
		case "q", "<C-c>":
			return nil
		}
		ui.Render()
	}
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
