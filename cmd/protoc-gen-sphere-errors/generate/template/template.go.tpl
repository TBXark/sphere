func (e {{.Name}}) Error() string {
    switch e {
    {{- range .Errors }}
    case {{.Name}}_{{.Value}}:
    {{ if .HasReason }}return "{{ .Reason }}";{{ else }}return "{{.Name}}:" + e.String(){{ end }}
    {{- end }}
    default:
        return "{{.Name}}:UNKNOWN_ERROR";
    }
}

func (e {{.Name}}) GetCode() int32 {
    switch e {
    {{- range .Errors }}
    case {{.Name}}_{{.Value}}:
        return {{.Code}};
    {{- end }}
    default:
        return 0;
    }
}

func (e {{.Name}}) GetStatus() int32 {
    switch e {
    {{- range .Errors }}
    case {{.Name}}_{{.Value}}:
        return {{.Status}};
    {{- end }}
    default:
        return 500;
    }
}

func (e {{.Name}}) GetMessage() string {
    switch e {
    {{- range .Errors }}
    case {{.Name}}_{{.Value}}:
        return "{{.Message}}";
    {{- end }}
    default:
        return "";
    }
}

func (e {{.Name}}) Join(errs ...error) error {
    allErrs := append(errs, e)
    msg := e.GetMessage()
    if msg == "" {
        msg = e.Error()
    }
    return statuserr.NewError(
        e.GetStatus(),
        e.GetCode(),
        msg,
        errors.Join(allErrs...),
    )
}

func (e {{.Name}}) JoinWithMessage(msg string, errs ...error) error {
    allErrs := append(errs, e)
	return statuserr.NewError(
	    e.GetStatus(),
	    e.GetCode(),
	    msg,
        errors.Join(allErrs...),
	)
}