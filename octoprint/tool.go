package octoprint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

const URITool = "/api/printer/tool"

// ToolCommand retrieves the current temperature data (actual, target and
// offset) plus optionally a (limited) history (actual, target, timestamp) for
// all of the printer’s available tools.
type ToolCommand struct {
	// History if true retrieve the temperature history.
	History bool
	// Limit limtis amount of returned history data points.
	Limit int
}

type ToolResponse toolResponse
type toolResponse struct {
	Current map[string]CurrentState `json:"current"`
	History []*History              `json:"history"`
}

func (r *ToolResponse) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	history := raw["history"]
	delete(raw, "history")
	b, _ = json.Marshal(map[string]interface{}{
		"current": raw,
		"history": history,
	})

	i := &toolResponse{}
	if err := json.Unmarshal(b, i); err != nil {
		return err
	}

	*r = ToolResponse(*i)
	return nil
}

// Do sends an API request and returns the API response.
func (cmd *ToolCommand) Do(c *Client) (*ToolResponse, error) {
	uri := fmt.Sprintf("%s?history=%t&limit=%d", URITool, cmd.History, cmd.Limit)
	b, err := c.doRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	r := &ToolResponse{}
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return r, err
}

// TargetCommand sets the given target temperature on the printer’s tools.
type TargetCommand struct {
	// Target temperature(s) to set, key must match the format tool{n} with n
	// being the tool’s index starting with 0.
	Target map[string]int `json:"target"`
}

// Do sends an API request and returns an error if any.
func (cmd *TargetCommand) Do(c *Client) error {
	b := bytes.NewBuffer(nil)
	if err := cmd.encode(b); err != nil {
		return err
	}

	_, err := c.doRequest("POST", URITool, b)
	return err
}

func (cmd *TargetCommand) encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(struct {
		Command string `json:"command"`
		TargetCommand
	}{
		Command:       "target",
		TargetCommand: *cmd,
	})
}

// OffsetCommand sets the given temperature offset on the printer’s tools.
type OffsetCommand struct {
	// Offset is offset(s) to set, key must match the format tool{n} with n
	// being the tool’s index starting with 0.
	Offsets map[string]int `json:"offsets"`
}

// Do sends an API request and returns an error if any.
func (cmd *OffsetCommand) Do(c *Client) error {
	b := bytes.NewBuffer(nil)
	if err := cmd.encode(b); err != nil {
		return err
	}

	_, err := c.doRequest("POST", URITool, b)
	return err
}

func (cmd *OffsetCommand) encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(struct {
		Command string `json:"command"`
		OffsetCommand
	}{
		Command:       "offset",
		OffsetCommand: *cmd,
	})
}

// ExtrudeCommand extrudes the given amount of filament from the currently
// selected tool.
type ExtrudeCommand struct {
	// Amount is the amount of filament to extrude in mm. May be negative to
	// retract.
	Amount int `json:"amount"`
}

// Do sends an API request and returns an error if any.
func (cmd *ExtrudeCommand) Do(c *Client) error {
	b := bytes.NewBuffer(nil)
	if err := cmd.encode(b); err != nil {
		return err
	}

	_, err := c.doRequest("POST", URITool, b)
	return err
}

func (cmd *ExtrudeCommand) encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(struct {
		Command string `json:"command"`
		ExtrudeCommand
	}{
		Command:        "extrude",
		ExtrudeCommand: *cmd,
	})
}

// SelectCommand selects the printer’s current tool.
type SelectCommand struct {
	// Tool to select, format tool{n} with n being the tool’s index starting
	// with 0.
	Tool string `json:"tool"`
}

// Do sends an API request and returns an error if any.
func (cmd *SelectCommand) Do(c *Client) error {
	b := bytes.NewBuffer(nil)
	if err := cmd.encode(b); err != nil {
		return err
	}

	_, err := c.doRequest("POST", URITool, b)
	return err
}

func (cmd *SelectCommand) encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(struct {
		Command string `json:"command"`
		SelectCommand
	}{
		Command:       "select",
		SelectCommand: *cmd,
	})
}

// FlowrateCommand changes the flow rate factor to apply to extrusion of the tool.
type FlowrateCommand struct {
	// Factor is the new factor, percentage as integer, between 75 and 125%.
	Factor string `json:"factor"`
}

// Do sends an API request and returns an error if any.
func (cmd *FlowrateCommand) Do(c *Client) error {
	b := bytes.NewBuffer(nil)
	if err := cmd.encode(b); err != nil {
		return err
	}

	_, err := c.doRequest("POST", URITool, b)
	return err
}

func (cmd *FlowrateCommand) encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(struct {
		Command string `json:"command"`
		FlowrateCommand
	}{
		Command:         "flowrate",
		FlowrateCommand: *cmd,
	})
}