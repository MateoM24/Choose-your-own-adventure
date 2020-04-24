package model

import (
	"fmt"
	"strings"
)

func ParseToStories(data map[string]map[string]interface{}) *Adventure {
	adv := newAdventure()
	for k, v := range data {
		story := parseToStory(v)
		adv.putStoryNode(k, story)
	}
	return adv
}

func parseToStory(data map[string]interface{}) StoryNode {
	var storyNode StoryNode
	for k, v := range data {
		switch k {
		case "title":
			storyNode.Title = v.(string)
		case "story":
			s := fmt.Sprintf("%v", v)
			s = strings.TrimPrefix(s, "[")
			s = strings.TrimSuffix(s, "]")
			storyNode.Story = s
		case "options":
			options := []Option{}
			for _, i := range v.([]interface{}) {
				row := i.(map[string]interface{})
				opt := Option{Arc: fmt.Sprintf("%v", row["arc"]), Text: fmt.Sprintf("%v", row["text"])}
				options = append(options, opt)
			}
			storyNode.Options = options
		}
	}
	return storyNode
}
