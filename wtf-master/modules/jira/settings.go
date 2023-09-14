package jira

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Jira"
)

type colors struct {
	rows struct {
		even string
		odd  string
	}
}

type Settings struct {
	colors
	common *cfg.Common

	apiKey                  string   `help:"Your Jira API key (or password for basic auth)."`
	domain                  string   `help:"Your Jira corporate domain."`
	email                   string   `help:"The email address associated with your Jira account (or username for basic auth)."`
	jql                     string   `help:"Custom JQL to be appended to the search query." values:"See Search Jira like a boss with JQL for details." optional:"true"`
	projects                []string `help:"An array of projects to get data from"`
	username                string   `help:"Your Jira username."`
	verifyServerCertificate bool     `help:"Determines whether or not the server’s certificate chain and host name are verified." values:"true or false" optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:                  ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_JIRA_API_KEY"))),
		domain:                  ymlConfig.UString("domain"),
		email:                   ymlConfig.UString("email"),
		jql:                     ymlConfig.UString("jql"),
		username:                ymlConfig.UString("username"),
		verifyServerCertificate: ymlConfig.UBool("verifyServerCertificate", true),
	}

	settings.colors.rows.even = ymlConfig.UString("colors.even", "lightblue")
	settings.colors.rows.odd = ymlConfig.UString("colors.odd", "white")

	settings.projects = settings.arrayifyProjects(ymlConfig, globalConfig)

	return &settings
}

/* -------------------- Unexported functions -------------------- */

// arrayifyProjects figures out if we're dealing with a single project or an array of projects
func (settings *Settings) arrayifyProjects(ymlConfig *config.Config, globalConfig *config.Config) []string {
	projects := []string{}

	// Single project
	project, err := ymlConfig.String("project")
	if err == nil {
		projects = append(projects, project)
		return projects
	}

	// Array of projects
	projectList := ymlConfig.UList("project")
	for _, projectName := range projectList {
		if project, ok := projectName.(string); ok {
			projects = append(projects, project)
		}
	}

	return projects
}