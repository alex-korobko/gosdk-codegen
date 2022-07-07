# gosdk-codegen

### Usage

Run `make all` command to generate go files from amazon specs. Before generation, latest changes are pulled from Amazon repo [https://github.com/amzn/selling-partner-api-models](https://github.com/amzn/selling-partner-api-models) Generated files stored at [amzn/selling-partner-api-go-sdk](amzn/selling-partner-api-go-sdk). 

Note that if Amazon has several specs code for each spec is generated in separated directoires but under the same package name.