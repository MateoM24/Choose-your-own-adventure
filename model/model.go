package model

type Option struct {
	Arc  string
	Text string
}

type StoryNode struct {
	Options []Option
	Story   string
	Title   string
}

type Adventure struct {
	nodes            map[string]StoryNode
	currentStoryNode StoryNode
}

func (adv *Adventure) Next(storyNodeTitle string) {
	adv.currentStoryNode = adv.nodes[storyNodeTitle]
}

func (adv *Adventure) start(startNodeName string) {
	adv.Next(startNodeName)
}

func (adv *Adventure) GetCurrentStoryNode() StoryNode {
	return adv.currentStoryNode
}

func (adv *Adventure) putStoryNode(name string, storyNode StoryNode) {
	adv.nodes[name] = storyNode
}

func newAdventure() *Adventure {
	adventure := Adventure{nodes: make(map[string]StoryNode)}
	return &adventure
}