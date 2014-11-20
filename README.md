# goramoph
## About
`goramoph` is a tool to transfer the play history of iTunes to BigQuery.
`goramoph` is a parody of `gramophone`.

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
