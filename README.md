gopocket
=============

Go package for consuming the PocketAPI (http://getpocket.com).  Requires a valid consumer key & access token from Pocket.

Working (w/ rate information)
_______
* Add
* Modify

In Progress
_______
* Retrieve

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