source = ["{{ .Values.Path }}"]
bundle_id = "{{ .Values.AppleBundleIDPrefix }}.{{ .Values.ProjectName }}"

sign {
    application_identity = "{{ .Values.AppleCertID }}"
}

dmg {
    output_path = "./dist/{{ .Values.ProjectName }}-v{{ .Values.Version }}-macos-{{ .Values.Arch }}.dmg"
    volume_name = "{{ .Values.ProjectName }}-{{ .Values.Arch }}"
}
