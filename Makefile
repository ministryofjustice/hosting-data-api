.PHONY: fetch-aws-billing-data
fetch-aws-billing-data:
	mkdir -p tmp/aws/billing-data
	aws s3 sync s3://moj-cur-reports/ tmp/aws/billing-data --exclude="*" --include="*-aws-billing-csv-*.csv"

.PHONY: aggregate-aws-billing-data
aggregate-aws-billing-data:
	csvstack tmp/aws/billing-data/*.csv > tmp/aws/aggregated-billing-data.csv

.PHONY: generate-aws-account-data
generate-aws-account-data:
	mkdir -p data/aws
	csvgrep -c RecordType -m AccountTotal tmp/aws/aggregated-billing-data.csv | csvcut -x -c LinkedAccountId,LinkedAccountName,BillingPeriodStartDate,BillingPeriodEndDate | csvjson -I | \
	jq '. | sort_by(.BillingPeriodStartDate) | group_by(.LinkedAccountId) | map({ id: .[0].LinkedAccountId } + { title: .[-1].LinkedAccountName, billing_start: .[0].BillingPeriodStartDate, billing_end: .[-1].BillingPeriodEndDate })' \
	> data/aws/accounts.json
