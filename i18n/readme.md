# i18n

* LocalizeConfig
* MustLocalize
* NewLocalizer
* Bundle
  * RegisterUnmarshalFunc
  * LoadMessageFile
    
To create a new localizer we need to pass it a bundle .

To create a new bundle we need to pass it the output object of language.Parse

The language package : "golang.org/x/text/language"

bundle needs to unmarshal a file , for this purpose
we must register a marshal function for bundle .

