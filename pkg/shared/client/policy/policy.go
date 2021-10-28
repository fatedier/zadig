package policy

import (
	"fmt"

	"github.com/koderover/zadig/pkg/setting"
	"github.com/koderover/zadig/pkg/tool/httpclient"
)

type Policy struct {
	Resource    string  `json:"resource"`
	Alias       string  `json:"alias"`
	Description string  `json:"description"`
	Rules       []*Rule `json:"rules"`
}

type Rule struct {
	Action      string        `json:"action"`
	Alias       string        `json:"alias"`
	Description string        `json:"description"`
	Rules       []*ActionRule `json:"rules"`
}

type ActionRule struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

func (c *Client) CreateOrUpdatePolicy(p *Policy) error {
	url := fmt.Sprintf("/policies/%s", p.Resource)

	_, err := c.Put(url, httpclient.SetBody(p))

	return err
}

type RoleBinding struct {
	Name   string
	User   string
	Role   setting.RoleType
	Public bool
}

func (c *Client) CreateRoleBinding(projectName string, roleBinding *RoleBinding) error {
	url := fmt.Sprintf("/rolebindings?projectName=%s", projectName)
	_, err := c.Post(url, httpclient.SetBody(roleBinding))
	return err
}

func (c *Client) DeleteRoleBinding(name string, projectName string) error {
	url := fmt.Sprintf("/rolebindings/%s?projectName=%s", name, projectName)
	_, err := c.Delete(url)
	return err
}

func (c *Client) CreateSystemRole(name string, role *Role) error {
	url := fmt.Sprintf("/system-roles/%s", name)
	_, err := c.Put(url, httpclient.SetBody(role))
	return err
}

func (c *Client) CreatePublicRole(name string, role *Role) error {
	url := fmt.Sprintf("/public-roles/%s", name)
	_, err := c.Put(url, httpclient.SetBody(role))
	return err
}

type Role struct {
	Name  string `json:"name"`
	Rules []*struct {
		Verbs     []string `json:"verbs"`
		Resources []string `json:"resources"`
		Kind      string   `json:"kind"`
	} `json:"rules"`
}
