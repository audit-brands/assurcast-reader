package storage

import "log"

// SeedDefaultSubscriptions adds the Assurcast feed bundle to a fresh database.
// No-op if any folders or feeds already exist — we never re-seed after first run,
// and we never disturb an existing user's subscription list.
func (s *Storage) SeedDefaultSubscriptions() {
	var count int
	err := s.db.QueryRow(
		`select (select count(*) from folders) + (select count(*) from feeds)`,
	).Scan(&count)
	if err != nil {
		log.Printf("seed: failed to check db state: %v", err)
		return
	}
	if count > 0 {
		return
	}

	folder := s.CreateFolder("Assurcast")
	if folder == nil {
		log.Print("seed: failed to create Assurcast folder, skipping feed seeds")
		return
	}

	seeds := []struct {
		title, link, feedLink string
	}{
		{"All Stories", "https://assurcast.com/", "https://assurcast.com/feed.rss"},
		{"Events", "https://assurcast.com/category/events", "https://assurcast.com/feed/category/events.rss"},
		{"IIA Standards", "https://assurcast.com/category/iia-standards", "https://assurcast.com/feed/category/iia-standards.rss"},
		{"Jobs", "https://assurcast.com/jobs", "https://assurcast.com/feed/jobs.rss"},
		{"News and Blogs", "https://assurcast.com/category/news", "https://assurcast.com/feed/category/news.rss"},
		{"Social and Media", "https://assurcast.com/category/social", "https://assurcast.com/feed/category/social.rss"},
		{"Tools and Tech", "https://assurcast.com/category/tools", "https://assurcast.com/feed/category/tools.rss"},
	}

	for _, seed := range seeds {
		if s.CreateFeed(seed.title, "", seed.link, seed.feedLink, &folder.Id) == nil {
			log.Printf("seed: failed to create feed %q", seed.title)
		}
	}

	log.Printf("seeded default Assurcast subscriptions: 1 folder, %d feeds", len(seeds))
}
