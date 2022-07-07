# gosdk-codegen

### Usage

Run `make all` command to generate go files from Amazon specs. Before generation, the latest changes are pulled from the Amazon repo [https://github.com/amzn/selling-partner-api-models](https://github.com/amzn/selling-partner-api-models). Generated files are stored at [amzn/selling-partner-api-go-sdk](amzn/selling-partner-api-go-sdk). 

Note that if Amazon has several specs code for each spec is generated in separated directories but under the same package name, as they not intdended to be used at once.