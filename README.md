gopocket
=============

Go package for consuming the PocketAPI (http://getpocket.com).  Requires a valid consumer key & access token from Pocket.

Add Usage
-------
     pocket := gopocket.Init(key, token)
     url := "http://an-interesting-article"
     title := "Won't be used if correctly parsed by Pocket, so it's just a backup."
     tags := []string{"An", "Array", "Of", "Tags"}

     response, rate, err := pocket.Add(url, title, tags)

Modify Usage
-------
     pocket := gopocket.Init(key, token)
     batch := gopocket.NewBatch()

     batch.Add("http://another-good-article", "", []string{"Article", "Google", "TV"})
     batch.Add("http://this-one-kinda-sucks", "", []string{"Terrible", "Bad Writing"})

     response, rate, err := pocket.Modify(batch)

Retrieve Usage
-------
     pocket := gopocket.Init(consumer_key, access_token)

     // all option method calls stack with previous options.
     opts := gopocket.NewOptions()
     t := time.Date(2013, 5, 19, 0, 0, 0, 0, time.UTC)

     opts.Since(t)                     // only return items since specified time...
     opts.Unfavorited()                // and are unfavorited...
     opts.State(gopocket.STATE_UNREAD) // and are unread.

     result, rate, err := pocket.Retrieve(opts)
     if err != nil {
          fmt.Println(err)
     } else {
          fmt.Println(result)
          fmt.Printf("\nCurrent Rate Info: %v\n", rate)
     }