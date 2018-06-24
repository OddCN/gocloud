package lambda

func preparedeleteserviceparams(params map[string]interface{}, DeleteFunction DeleteFunction, Region string) {
	if Region != "" {
		params["Region"] = Region
	}

	if DeleteFunction.FunctionName != "" {
		params["FunctionName"] = DeleteFunction.FunctionName
	}

	if DeleteFunction.Qualifier != "" {
		params["Qualifier"] = DeleteFunction.Qualifier
	}

}
