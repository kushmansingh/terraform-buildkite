package buildkite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePipeline() *schema.Resource {
	return &schema.Resource{
		Create: CreatePipeline,
		Read:   ReadPipeline,
		Update: UpdatePipeline,
		Delete: DeletePipeline,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"web_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"builds_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"badge_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"branch_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_branch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "master",
			},
			"env": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"webhook_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"skip_queued_branch_builds": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"skip_queued_branch_builds_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cancel_running_branch_builds": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cancel_running_branch_builds_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"label": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"command": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"env": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"timeout_in_minutes": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"agent_query_rules": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"artifact_paths": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"branch_configuration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"concurrency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"parallelism": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"bitbucket_settings": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"github_settings"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"build_pull_requests": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"pull_request_branch_filter_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"pull_request_branch_filter_configuration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"skip_pull_request_builds_for_existing_commits": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"build_tags": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"publish_commit_status": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"publish_commit_status_per_step": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"github_settings": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"bitbucket_settings"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trigger_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default: "none",
						},
						"build_pull_requests": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"pull_request_branch_filter_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"pull_request_branch_filter_configuration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"skip_pull_request_builds_for_existing_commits": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"separate_pull_request_statuses": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"skip_builds_for_existing_commits": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"filter_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"build_pull_request_forks": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"prefix_pull_request_fork_branch_names": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"build_tags": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"publish_commit_status": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"publish_commit_status_per_step": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"publish_blocked_as_pending": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

type Pipeline struct {
	Id                              string                 `json:"id,omitempty"`
	Environment                     envMap                 `json:"env,omitempty"`
	Slug                            string                 `json:"slug,omitempty"`
	WebURL                          string                 `json:"web_url,omitempty"`
	BuildsURL                       string                 `json:"builds_url,omitempty"`
	Url                             string                 `json:"url,omitempty"`
	DefaultBranch                   string                 `json:"default_branch,omitempty"`
	BadgeURL                        string                 `json:"badge_url,omitempty"`
	CreatedAt                       string                 `json:"created_at,omitempty"`
	Repository                      string                 `json:"repository,omitempty"`
	Name                            string                 `json:"name,omitempty"`
	Description                     string                 `json:"description,omitempty"`
	BranchConfiguration             string                 `json:"branch_configuration,omitempty"`
	SkipQueuedBranchBuilds          bool                   `json:"skip_queued_branch_builds,omitempty"`
	SkipQueuedBranchBuildsFilter    string                 `json:"skip_queued_branch_builds_filter,omitempty"`
	CancelRunningBranchBuilds       bool                   `json:"cancel_running_branch_builds,omitempty"`
	CancelRunningBranchBuildsFilter string                 `json:"cancel_running_branch_builds_filter,omitempty"`
	Provider                        repositoryProvider     `json:"provider,omitempty"`
	ProviderSettings                map[string]interface{} `json:"provider_settings,omitempty"`
	Steps                           []Step                 `json:"steps"`
}

type envMap map[string]string

func (e *envMap) UnmarshalJSON(data []byte) error {
	env := map[string]interface{}{}

	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	if err := d.Decode(&env); err != nil {
		return err
	}
	m := make(envMap)
	for k, v := range env {
		switch val := v.(type) {
		case json.Number:
			m[k] = val.String()
		case bool:
			m[k] = strconv.FormatBool(val)
		case string:
			m[k] = val
		}
	}
	*e = m
	return nil
}

type repositoryProvider struct {
	RepositoryProviderId string
	Settings             map[string]interface{}
	WebhookURL           string
}

var providerSettingsExcluded = [...]string{"repository", "account", "commit_status_404s", "commit_status_error"}

func (p repositoryProvider) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"id":          p.RepositoryProviderId,
		"settings":    p.Settings,
		"webhook_url": p.WebhookURL,
	}
	return json.Marshal(m)
}

func (p *repositoryProvider) UnmarshalJSON(data []byte) error {
	var provider map[string]interface{}

	if err := json.Unmarshal(data, &provider); err != nil {
		return err
	}

	p.RepositoryProviderId = provider["id"].(string)
	webhook, ok := provider["webhook_url"]
	if ok {
		p.WebhookURL = webhook.(string)
	}

	settings := provider["settings"].(map[string]interface{})

	for _, k := range providerSettingsExcluded {
		delete(settings, k)
	}

	p.Settings = settings

	return nil
}

type Step struct {
	Type                string            `json:"type"`
	Name                string            `json:"name,omitempty"`
	Label               string            `json:"label,omitempty"`
	Command             string            `json:"command,omitempty"`
	Environment         map[string]string `json:"env,omitempty"`
	TimeoutInMinutes    int               `json:"timeout_in_minutes,omitempty"`
	AgentQueryRules     []string          `json:"agent_query_rules,omitempty"`
	BranchConfiguration string            `json:"branch_configuration,omitempty"`
	ArtifactPaths       string            `json:"artifact_paths,omitempty"`
	Concurrency         int               `json:"concurrency,omitempty"`
	Parallelism         int               `json:"parallelism,omitempty"`
}

func CreatePipeline(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[TRACE] CreatePipeline")

	client := meta.(*Client)

	req := preparePipelineRequestPayload(d)
	res := &Pipeline{}

	err := client.Post([]string{"pipelines"}, req, res)
	if err != nil {
		return err
	}

	return updatePipelineFromAPI(d, res)
}

func ReadPipeline(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[TRACE] ReadPipeline")

	client := meta.(*Client)
	slug := d.Id()

	res := &Pipeline{}

	err := client.Get([]string{"pipelines", slug}, res)
	if err != nil {
		if _, ok := err.(*notFound); ok {
			d.SetId("")
			return nil
		}
		return err
	}

	return updatePipelineFromAPI(d, res)
}

func UpdatePipeline(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[TRACE] UpdatePipeline")

	client := meta.(*Client)
	slug := d.Id()

	req := preparePipelineRequestPayload(d)
	res := &Pipeline{}

	err := client.Patch([]string{"pipelines", slug}, req, res)
	if err != nil {
		return err
	}

	return updatePipelineFromAPI(d, res)
}

func DeletePipeline(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[TRACE] DeletePipeline")

	client := meta.(*Client)

	slug := d.Id()

	return client.Delete([]string{"pipelines", slug})
}

func updatePipelineFromAPI(d *schema.ResourceData, p *Pipeline) error {
	d.SetId(p.Slug)
	log.Printf("[INFO] buildkite: Pipeline ID: %s", d.Id())

	d.Set("env", p.Environment)
	d.Set("name", p.Name)
	d.Set("description", p.Description)
	d.Set("repository", p.Repository)
	d.Set("web_url", p.WebURL)
	d.Set("slug", p.Slug)
	d.Set("builds_url", p.BuildsURL)
	d.Set("branch_configuration", p.BranchConfiguration)
	d.Set("default_branch", p.DefaultBranch)
	d.Set("skip_queued_branch_builds", p.SkipQueuedBranchBuilds)
	d.Set("skip_queued_branch_builds_filter", p.SkipQueuedBranchBuildsFilter)
	d.Set("cancel_running_branch_builds", p.CancelRunningBranchBuilds)
	d.Set("cancel_running_branch_builds_filter", p.CancelRunningBranchBuildsFilter)

	stepMap := make([]interface{}, len(p.Steps))
	for i, element := range p.Steps {
		stepMap[i] = map[string]interface{}{
			"type":                 element.Type,
			"label":                element.Label,
			"name":                 element.Name,
			"command":              element.Command,
			"env":                  element.Environment,
			"agent_query_rules":    element.AgentQueryRules,
			"branch_configuration": element.BranchConfiguration,
			"artifact_paths":       element.ArtifactPaths,
			"concurrency":          element.Concurrency,
			"parallelism":          element.Parallelism,
			"timeout_in_minutes":   element.TimeoutInMinutes,
		}
	}
	if err := d.Set("step", stepMap); err != nil {
		return err
	}

	emptySettings := make([]interface{}, 0)
	d.Set("github_settings", emptySettings)
	d.Set("bitbucket_settings", emptySettings)

	log.Printf("[INFO] buildkite: RepositoryProviderId: %s", p.Provider.RepositoryProviderId)

	switch p.Provider.RepositoryProviderId {
	case "github":
		d.Set("webhook_url", p.Provider.WebhookURL)

		log.Printf("[DEBUG] buildkite: Provider.Settings in github: %+v", p.Provider.Settings)
		if err := d.Set("github_settings", []map[string]interface{}{p.Provider.Settings}); err != nil {
			return err
		}

	case "bitbucket":
		d.Set("webhook_url", p.Provider.WebhookURL)

		log.Printf("[DEBUG] buildkite: Provider.Settings in bitbucket: %+v", p.Provider.Settings)
		if err := d.Set("bitbucket_settings", []map[string]interface{}{p.Provider.Settings}); err != nil {
			return err
		}

	case "gitlab":
		d.Set("webhook_url", p.Provider.WebhookURL)

	case "beanstalk":
		d.Set("webhook_url", p.Provider.WebhookURL)

	default: // unknown, noop
	}

	return nil
}

func preparePipelineRequestPayload(d *schema.ResourceData) *Pipeline {
	req := &Pipeline{}

	req.Name = d.Get("name").(string)
	req.DefaultBranch = d.Get("default_branch").(string)
	req.Description = d.Get("description").(string)
	req.Slug = d.Get("slug").(string)
	req.Repository = d.Get("repository").(string)
	req.BranchConfiguration = d.Get("branch_configuration").(string)
	req.Environment = d.Get("environment").(envMap)
	req.SkipQueuedBranchBuilds = d.Get("skip_queued_branch_builds").(bool)
	req.SkipQueuedBranchBuildsFilter = d.Get("skip_queued_branch_builds_filter").(string)
	req.CancelRunningBranchBuilds = d.Get("cancel_running_branch_builds").(bool)
	req.CancelRunningBranchBuildsFilter = d.Get("cancel_running_branch_builds_filter").(string)

	stepsI := d.Get("step").([]interface{})
	req.Steps = make([]Step, len(stepsI))

	for i, stepI := range stepsI {
		stepM := stepI.(map[string]interface{})
		req.Steps[i] = Step{
			Type:                stepM["type"].(string),
			Label:               stepM["label"].(string),
			Name:                stepM["name"].(string),
			Command:             stepM["command"].(string),
			Environment:         map[string]string{},
			AgentQueryRules:     make([]string, len(stepM["agent_query_rules"].([]interface{}))),
			BranchConfiguration: stepM["branch_configuration"].(string),
			ArtifactPaths:       stepM["artifact_paths"].(string),
			Concurrency:         stepM["concurrency"].(int),
			Parallelism:         stepM["parallelism"].(int),
			TimeoutInMinutes:    stepM["timeout_in_minutes"].(int),
		}

		for k, vI := range stepM["env"].(map[string]interface{}) {
			req.Steps[i].Environment[k] = vI.(string)
		}

		for j, vI := range stepM["agent_query_rules"].([]interface{}) {
			req.Steps[i].AgentQueryRules[j] = vI.(string)
		}
	}

	githubSettings := d.Get("github_settings").([]interface{})
	bitbucketSettings := d.Get("bitbucket_settings").([]interface{})
	req.Provider = repositoryProvider{}
	var stateSettings map[string]interface{}
	var stateName string
	if len(githubSettings) > 0 {
		req.Provider.RepositoryProviderId = "github"
		stateSettings = githubSettings[0].(map[string]interface{})
		stateName = "github_settings"
	} else if len(bitbucketSettings) > 0 {
		req.Provider.RepositoryProviderId = "bitbucket"
		stateSettings = bitbucketSettings[0].(map[string]interface{})
		stateName = "bitbucket_settings"
	}

	req.Provider.Settings = map[string]interface{}{}
	// If we have either settings
	if stateName != "" {

		for k, vI := range stateSettings {
			if _, ok := d.GetOk(fmt.Sprintf("%s.0.%s", stateName, k)); ok {
				req.Provider.Settings[k] = vI
			}
		}
	}

	return req
}
