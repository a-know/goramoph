# goramoph
## About
`goramoph` is a tool to load the play history of iTunes to BigQuery.


## Prepare
1. Download the `Google Cloud SDK` and dispense the authentication.
    * https://cloud.google.com/sdk/
        * FYI http://qiita.com/yuko/items/1c4ee5b081c5b6a3ac8a
    * if you use `zsh`, see also http://tsukaby.com/tech_blog/archives/482
    * Following the approval authority is the minimum required
        * `Google Cloud Storage`
        * `Google BigQuery`
2. Create Project for goramoph
	* Access https://console.developers.google.com/ and signin.
	* Click [Create Project] and create Project.
		* Please specify `goramoph` in `Project Name`.
		* Please specify unique string in `Project ID`. (ex. a-know-goramoph)
			* You have to enable billing settings.
3. Project Setting
	* `$ gcloud config set project <project-id>`
		* ex. `$ gcloud config set project a-know-goramoph`


## How to use
`$ go run goramoph.go [path to iTunes Music Library.xml]`


## Goramoph's behavior
1. Parse `iTunes Music Library.xml`.
2. Parsing result is export to `./csv` as csv file.
    * Csv file name is `./csv/<last-modify-date>.csv`
3. Upload csv file to `Google Cloud Storage`.
    * Bucket name is `<project-id>-csv`
4. Load csv file contents to `Google BigQuery` dataset and table.
    * Dataset name is `<project-id>_ds`
    * table name is `<last-modify-date>`
5. Remove uploaded csv file.


## Notes
* `Google Cloud Storage` and `Google BigQuery` are require billing setting. 
* Absolutely I do not know anything about your billing amounts.


## License
This software is released under the MIT License, see LICENSE.txt.
