package controllers

import (
	"bufio"
	"os"
	"strings"

	"github.com/vuonglequoc/openvpn-web-ui/models"

	"github.com/beego/beego/v2/core/logs"
)

type LogsController struct {
	BaseController
}

func (c *LogsController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
}

func (c *LogsController) Get() {
	c.TplName = "logs.html"
	c.Data["breadcrumbs"] = &BreadCrumbs{
		Title: "Logs",
	}

	settings := models.Settings{Profile: "default"}
	settings.Read("Profile")

	if err := settings.Read("OVConfigPath"); err != nil {
		logs.Error(err)
		return
	}

	// run_logs
	fName := settings.OVConfigPath + "log/openvpn.log"
	file, err := os.Open(fName)
	if err != nil {
		logs.Error(err)
	} else {
		scanner := bufio.NewScanner(file)
		var run_logs []string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Index(line, " MANAGEMENT: ") == -1 {
				run_logs = append(run_logs, strings.Trim(line, "\t"))
			}
		}
		start := len(run_logs) - 200
		if start < 0 {
			start = 0
		}
		c.Data["logs"] = reverse(run_logs[start:])
	}
	defer file.Close()

	// status_logs
	fName = settings.OVConfigPath + "log/openvpn-status.log"
	file, err = os.Open(fName)
	if err != nil {
		logs.Error(err)
	} else {
		scanner := bufio.NewScanner(file)
		var status_logs []string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Index(line, " MANAGEMENT: ") == -1 {
				status_logs = append(status_logs, strings.Trim(line, "\t"))
			}
		}
		start := len(status_logs) - 200
		if start < 0 {
			start = 0
		}
		c.Data["status_logs"] = reverse(status_logs[start:])
	}
	defer file.Close()
}

func reverse(lines []string) []string {
	for i := 0; i < len(lines)/2; i++ {
		j := len(lines) - i - 1
		lines[i], lines[j] = lines[j], lines[i]
	}
	return lines
}
