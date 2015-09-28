package hn

const (
	STORY_TYPE_NEW  = "new"
	STORY_TYPE_TOP  = "top"
	STORY_TYPE_JOB  = "job"
	STORY_TYPE_ASK  = "ask"
	STORY_TYPE_SHOW = "show"
)

// LiveService communicates with the news
// related endpoints in the Hacker News API
type LiveService interface {
	GetStories(string) ([]int, error)

	MaxItem() (int, error)
	Updates() (*Updates, error)
}

// liveService implements LiveService.
type liveService struct {
	client *Client
}

// Updates contains the latest updated items and profiles
type Updates struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}

func (s *liveService) GetStories(storyType string) ([]int, error) {
	req, err := s.client.NewRequest(s.storiesPath(storyType))
	if err != nil {
		return nil, err
	}

	var value []int
	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *Client) TopStories() ([]int, error) {
	return c.Live.GetStories(STORY_TYPE_TOP)
}

func (c *Client) NewStories() ([]int, error) {
	return c.Live.GetStories(STORY_TYPE_NEW)
}

func (c *Client) AskStories() ([]int, error) {
	return c.Live.GetStories(STORY_TYPE_ASK)
}

func (c *Client) JobStories() ([]int, error) {
	return c.Live.GetStories(STORY_TYPE_JOB)
}

func (c *Client) ShowStories() ([]int, error) {
	return c.Live.GetStories(STORY_TYPE_SHOW)
}

func (s *liveService) storiesPath(storyType string) string {
	return storyType + "stories.json"
}

// MaxItem is a convenience method proxying Live.MaxItem
func (c *Client) MaxItem() (int, error) {
	return c.Live.MaxItem()
}

// MaxItem retrieves the current largest item id
func (s *liveService) MaxItem() (int, error) {
	req, err := s.client.NewRequest(s.maxItemPath())
	if err != nil {
		return 0, err
	}

	var value int
	_, err = s.client.Do(req, &value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *liveService) maxItemPath() string {
	return "maxitem.json"
}

// Updates is a convenience method proxying Live.Updates
func (c *Client) Updates() (*Updates, error) {
	return c.Live.Updates()
}

// Updates retrieves the current largest item id
func (s *liveService) Updates() (*Updates, error) {
	req, err := s.client.NewRequest(s.updatesPath())
	if err != nil {
		return nil, err
	}

	var value Updates
	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (s *liveService) updatesPath() string {
	return "updates.json"
}
