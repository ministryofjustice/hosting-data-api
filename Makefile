.PHONY: fetch-aws-billing-data
fetch-aws-billing-data:
	mkdir -p tmp/aws
	aws s3 sync s3://moj-cur-reports/ tmp/aws --exclude="*" --include="*-aws-billing-csv-*.csv"
