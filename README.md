gopocket
=============

Go package for consuming the PocketAPI (http://getpocket.com).  Requires a valid consumer key & access token from Pocket.

Working
_______
* Add (w/ rate information)

In Progress
_______
* Modify
* Retrieve

Usage
-------
     pocket := gopocket.Init(key, token)
     url := "http://an-interesting-article"
     title := "Won't be used if correctly parsed by Pocket, so it's just a backup."
     tags = []string{"An", "Array", "Of", "Tags"}

     response, rate, err := pocket.Add(url, title, tags)