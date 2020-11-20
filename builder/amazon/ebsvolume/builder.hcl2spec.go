// Code generated by "mapstructure-to-hcl2 -type Config,BlockDevice"; DO NOT EDIT.
package ebsvolume

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/builder/amazon/common"
	"github.com/hashicorp/packer/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

// FlatBlockDevice is an auto-generated flat version of BlockDevice.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatBlockDevice struct {
	DeleteOnTermination *bool                 `mapstructure:"delete_on_termination" required:"false" cty:"delete_on_termination" hcl:"delete_on_termination"`
	DeviceName          *string               `mapstructure:"device_name" required:"false" cty:"device_name" hcl:"device_name"`
	Encrypted           *bool                 `mapstructure:"encrypted" required:"false" cty:"encrypted" hcl:"encrypted"`
	IOPS                *int64                `mapstructure:"iops" required:"false" cty:"iops" hcl:"iops"`
	NoDevice            *bool                 `mapstructure:"no_device" required:"false" cty:"no_device" hcl:"no_device"`
	SnapshotId          *string               `mapstructure:"snapshot_id" required:"false" cty:"snapshot_id" hcl:"snapshot_id"`
	VirtualName         *string               `mapstructure:"virtual_name" required:"false" cty:"virtual_name" hcl:"virtual_name"`
	VolumeType          *string               `mapstructure:"volume_type" required:"false" cty:"volume_type" hcl:"volume_type"`
	VolumeSize          *int64                `mapstructure:"volume_size" required:"false" cty:"volume_size" hcl:"volume_size"`
	KmsKeyId            *string               `mapstructure:"kms_key_id" required:"false" cty:"kms_key_id" hcl:"kms_key_id"`
	Tags                map[string]string     `mapstructure:"tags" required:"false" cty:"tags" hcl:"tags"`
	Tag                 []config.FlatKeyValue `mapstructure:"tag" required:"false" cty:"tag" hcl:"tag"`
}

// FlatMapstructure returns a new FlatBlockDevice.
// FlatBlockDevice is an auto-generated flat version of BlockDevice.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*BlockDevice) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatBlockDevice)
}

// HCL2Spec returns the hcl spec of a BlockDevice.
// This spec is used by HCL to read the fields of BlockDevice.
// The decoded values from this spec will then be applied to a FlatBlockDevice.
func (*FlatBlockDevice) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"delete_on_termination": &hcldec.AttrSpec{Name: "delete_on_termination", Type: cty.Bool, Required: false},
		"device_name":           &hcldec.AttrSpec{Name: "device_name", Type: cty.String, Required: false},
		"encrypted":             &hcldec.AttrSpec{Name: "encrypted", Type: cty.Bool, Required: false},
		"iops":                  &hcldec.AttrSpec{Name: "iops", Type: cty.Number, Required: false},
		"no_device":             &hcldec.AttrSpec{Name: "no_device", Type: cty.Bool, Required: false},
		"snapshot_id":           &hcldec.AttrSpec{Name: "snapshot_id", Type: cty.String, Required: false},
		"virtual_name":          &hcldec.AttrSpec{Name: "virtual_name", Type: cty.String, Required: false},
		"volume_type":           &hcldec.AttrSpec{Name: "volume_type", Type: cty.String, Required: false},
		"volume_size":           &hcldec.AttrSpec{Name: "volume_size", Type: cty.Number, Required: false},
		"kms_key_id":            &hcldec.AttrSpec{Name: "kms_key_id", Type: cty.String, Required: false},
		"tags":                  &hcldec.AttrSpec{Name: "tags", Type: cty.Map(cty.String), Required: false},
		"tag":                   &hcldec.BlockListSpec{TypeName: "tag", Nested: hcldec.ObjectSpec((*config.FlatKeyValue)(nil).HCL2Spec())},
	}
	return s
}

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName                           *string                                `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType                         *string                                `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion                         *string                                `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug                               *bool                                  `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce                               *bool                                  `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError                             *string                                `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars                            map[string]string                      `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars                       []string                               `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	AccessKey                                 *string                                `mapstructure:"access_key" required:"true" cty:"access_key" hcl:"access_key"`
	AssumeRole                                *common.FlatAssumeRoleConfig           `mapstructure:"assume_role" required:"false" cty:"assume_role" hcl:"assume_role"`
	CustomEndpointEc2                         *string                                `mapstructure:"custom_endpoint_ec2" required:"false" cty:"custom_endpoint_ec2" hcl:"custom_endpoint_ec2"`
	CredsFilename                             *string                                `mapstructure:"shared_credentials_file" required:"false" cty:"shared_credentials_file" hcl:"shared_credentials_file"`
	DecodeAuthZMessages                       *bool                                  `mapstructure:"decode_authorization_messages" required:"false" cty:"decode_authorization_messages" hcl:"decode_authorization_messages"`
	InsecureSkipTLSVerify                     *bool                                  `mapstructure:"insecure_skip_tls_verify" required:"false" cty:"insecure_skip_tls_verify" hcl:"insecure_skip_tls_verify"`
	MaxRetries                                *int                                   `mapstructure:"max_retries" required:"false" cty:"max_retries" hcl:"max_retries"`
	MFACode                                   *string                                `mapstructure:"mfa_code" required:"false" cty:"mfa_code" hcl:"mfa_code"`
	ProfileName                               *string                                `mapstructure:"profile" required:"false" cty:"profile" hcl:"profile"`
	RawRegion                                 *string                                `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
	SecretKey                                 *string                                `mapstructure:"secret_key" required:"true" cty:"secret_key" hcl:"secret_key"`
	SkipMetadataApiCheck                      *bool                                  `mapstructure:"skip_metadata_api_check" cty:"skip_metadata_api_check" hcl:"skip_metadata_api_check"`
	SkipCredsValidation                       *bool                                  `mapstructure:"skip_credential_validation" cty:"skip_credential_validation" hcl:"skip_credential_validation"`
	Token                                     *string                                `mapstructure:"token" required:"false" cty:"token" hcl:"token"`
	VaultAWSEngine                            *common.FlatVaultAWSEngineOptions      `mapstructure:"vault_aws_engine" required:"false" cty:"vault_aws_engine" hcl:"vault_aws_engine"`
	PollingConfig                             *common.FlatAWSPollingConfig           `mapstructure:"aws_polling" required:"false" cty:"aws_polling" hcl:"aws_polling"`
	AssociatePublicIpAddress                  *bool                                  `mapstructure:"associate_public_ip_address" required:"false" cty:"associate_public_ip_address" hcl:"associate_public_ip_address"`
	AvailabilityZone                          *string                                `mapstructure:"availability_zone" required:"false" cty:"availability_zone" hcl:"availability_zone"`
	BlockDurationMinutes                      *int64                                 `mapstructure:"block_duration_minutes" required:"false" cty:"block_duration_minutes" hcl:"block_duration_minutes"`
	DisableStopInstance                       *bool                                  `mapstructure:"disable_stop_instance" required:"false" cty:"disable_stop_instance" hcl:"disable_stop_instance"`
	EbsOptimized                              *bool                                  `mapstructure:"ebs_optimized" required:"false" cty:"ebs_optimized" hcl:"ebs_optimized"`
	EnableT2Unlimited                         *bool                                  `mapstructure:"enable_t2_unlimited" required:"false" cty:"enable_t2_unlimited" hcl:"enable_t2_unlimited"`
	IamInstanceProfile                        *string                                `mapstructure:"iam_instance_profile" required:"false" cty:"iam_instance_profile" hcl:"iam_instance_profile"`
	SkipProfileValidation                     *bool                                  `mapstructure:"skip_profile_validation" required:"false" cty:"skip_profile_validation" hcl:"skip_profile_validation"`
	TemporaryIamInstanceProfilePolicyDocument *common.FlatPolicyDocument             `mapstructure:"temporary_iam_instance_profile_policy_document" required:"false" cty:"temporary_iam_instance_profile_policy_document" hcl:"temporary_iam_instance_profile_policy_document"`
	InstanceInitiatedShutdownBehavior         *string                                `mapstructure:"shutdown_behavior" required:"false" cty:"shutdown_behavior" hcl:"shutdown_behavior"`
	InstanceType                              *string                                `mapstructure:"instance_type" required:"true" cty:"instance_type" hcl:"instance_type"`
	SecurityGroupFilter                       *common.FlatSecurityGroupFilterOptions `mapstructure:"security_group_filter" required:"false" cty:"security_group_filter" hcl:"security_group_filter"`
	RunTags                                   map[string]string                      `mapstructure:"run_tags" required:"false" cty:"run_tags" hcl:"run_tags"`
	RunTag                                    []config.FlatKeyValue                  `mapstructure:"run_tag" required:"false" cty:"run_tag" hcl:"run_tag"`
	SecurityGroupId                           *string                                `mapstructure:"security_group_id" required:"false" cty:"security_group_id" hcl:"security_group_id"`
	SecurityGroupIds                          []string                               `mapstructure:"security_group_ids" required:"false" cty:"security_group_ids" hcl:"security_group_ids"`
	SourceAmi                                 *string                                `mapstructure:"source_ami" required:"true" cty:"source_ami" hcl:"source_ami"`
	SourceAmiFilter                           *common.FlatAmiFilterOptions           `mapstructure:"source_ami_filter" required:"false" cty:"source_ami_filter" hcl:"source_ami_filter"`
	SpotInstanceTypes                         []string                               `mapstructure:"spot_instance_types" required:"false" cty:"spot_instance_types" hcl:"spot_instance_types"`
	SpotPrice                                 *string                                `mapstructure:"spot_price" required:"false" cty:"spot_price" hcl:"spot_price"`
	SpotPriceAutoProduct                      *string                                `mapstructure:"spot_price_auto_product" required:"false" undocumented:"true" cty:"spot_price_auto_product" hcl:"spot_price_auto_product"`
	SpotTags                                  map[string]string                      `mapstructure:"spot_tags" required:"false" cty:"spot_tags" hcl:"spot_tags"`
	SpotTag                                   []config.FlatKeyValue                  `mapstructure:"spot_tag" required:"false" cty:"spot_tag" hcl:"spot_tag"`
	SubnetFilter                              *common.FlatSubnetFilterOptions        `mapstructure:"subnet_filter" required:"false" cty:"subnet_filter" hcl:"subnet_filter"`
	SubnetId                                  *string                                `mapstructure:"subnet_id" required:"false" cty:"subnet_id" hcl:"subnet_id"`
	Tenancy                                   *string                                `mapstructure:"tenancy" required:"false" cty:"tenancy" hcl:"tenancy"`
	TemporarySGSourceCidrs                    []string                               `mapstructure:"temporary_security_group_source_cidrs" required:"false" cty:"temporary_security_group_source_cidrs" hcl:"temporary_security_group_source_cidrs"`
	UserData                                  *string                                `mapstructure:"user_data" required:"false" cty:"user_data" hcl:"user_data"`
	UserDataFile                              *string                                `mapstructure:"user_data_file" required:"false" cty:"user_data_file" hcl:"user_data_file"`
	VpcFilter                                 *common.FlatVpcFilterOptions           `mapstructure:"vpc_filter" required:"false" cty:"vpc_filter" hcl:"vpc_filter"`
	VpcId                                     *string                                `mapstructure:"vpc_id" required:"false" cty:"vpc_id" hcl:"vpc_id"`
	WindowsPasswordTimeout                    *string                                `mapstructure:"windows_password_timeout" required:"false" cty:"windows_password_timeout" hcl:"windows_password_timeout"`
	Type                                      *string                                `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect                        *string                                `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                                   *string                                `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                                   *int                                   `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername                               *string                                `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword                               *string                                `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName                            *string                                `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName                   *string                                `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType                   *string                                `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits                   *int                                   `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                                []string                               `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys                    *bool                                  `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos                               []string                               `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile                         *string                                `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile                        *string                                `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                                    *bool                                  `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                                *string                                `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout                            *string                                `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth                              *bool                                  `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding                 *bool                                  `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts                      *int                                   `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost                            *string                                `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort                            *int                                   `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth                       *bool                                  `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername                        *string                                `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword                        *string                                `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive                     *bool                                  `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile                  *string                                `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile                 *string                                `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod                     *string                                `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost                              *string                                `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort                              *int                                   `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername                          *string                                `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword                          *string                                `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval                      *string                                `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout                       *string                                `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels                          []string                               `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels                           []string                               `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey                              []byte                                 `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey                             []byte                                 `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                                 *string                                `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword                             *string                                `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                                 *string                                `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy                              *bool                                  `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                                 *int                                   `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout                              *string                                `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL                               *bool                                  `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure                             *bool                                  `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM                              *bool                                  `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	SSHInterface                              *string                                `mapstructure:"ssh_interface" cty:"ssh_interface" hcl:"ssh_interface"`
	PauseBeforeSSM                            *string                                `mapstructure:"pause_before_ssm" cty:"pause_before_ssm" hcl:"pause_before_ssm"`
	SessionManagerPort                        *int                                   `mapstructure:"session_manager_port" cty:"session_manager_port" hcl:"session_manager_port"`
	AMIENASupport                             *bool                                  `mapstructure:"ena_support" required:"false" cty:"ena_support" hcl:"ena_support"`
	AMISriovNetSupport                        *bool                                  `mapstructure:"sriov_support" required:"false" cty:"sriov_support" hcl:"sriov_support"`
	VolumeMappings                            []FlatBlockDevice                      `mapstructure:"ebs_volumes" required:"false" cty:"ebs_volumes" hcl:"ebs_volumes"`
	VolumeRunTags                             map[string]string                      `mapstructure:"run_volume_tags" cty:"run_volume_tags" hcl:"run_volume_tags"`
	VolumeRunTag                              []config.FlatKeyValue                  `mapstructure:"run_volume_tag" cty:"run_volume_tag" hcl:"run_volume_tag"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":             &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":           &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":           &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                  &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                  &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":               &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":         &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":    &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"access_key":                    &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"assume_role":                   &hcldec.BlockSpec{TypeName: "assume_role", Nested: hcldec.ObjectSpec((*common.FlatAssumeRoleConfig)(nil).HCL2Spec())},
		"custom_endpoint_ec2":           &hcldec.AttrSpec{Name: "custom_endpoint_ec2", Type: cty.String, Required: false},
		"shared_credentials_file":       &hcldec.AttrSpec{Name: "shared_credentials_file", Type: cty.String, Required: false},
		"decode_authorization_messages": &hcldec.AttrSpec{Name: "decode_authorization_messages", Type: cty.Bool, Required: false},
		"insecure_skip_tls_verify":      &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"max_retries":                   &hcldec.AttrSpec{Name: "max_retries", Type: cty.Number, Required: false},
		"mfa_code":                      &hcldec.AttrSpec{Name: "mfa_code", Type: cty.String, Required: false},
		"profile":                       &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"region":                        &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"secret_key":                    &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"skip_metadata_api_check":       &hcldec.AttrSpec{Name: "skip_metadata_api_check", Type: cty.Bool, Required: false},
		"skip_credential_validation":    &hcldec.AttrSpec{Name: "skip_credential_validation", Type: cty.Bool, Required: false},
		"token":                         &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"vault_aws_engine":              &hcldec.BlockSpec{TypeName: "vault_aws_engine", Nested: hcldec.ObjectSpec((*common.FlatVaultAWSEngineOptions)(nil).HCL2Spec())},
		"aws_polling":                   &hcldec.BlockSpec{TypeName: "aws_polling", Nested: hcldec.ObjectSpec((*common.FlatAWSPollingConfig)(nil).HCL2Spec())},
		"associate_public_ip_address":   &hcldec.AttrSpec{Name: "associate_public_ip_address", Type: cty.Bool, Required: false},
		"availability_zone":             &hcldec.AttrSpec{Name: "availability_zone", Type: cty.String, Required: false},
		"block_duration_minutes":        &hcldec.AttrSpec{Name: "block_duration_minutes", Type: cty.Number, Required: false},
		"disable_stop_instance":         &hcldec.AttrSpec{Name: "disable_stop_instance", Type: cty.Bool, Required: false},
		"ebs_optimized":                 &hcldec.AttrSpec{Name: "ebs_optimized", Type: cty.Bool, Required: false},
		"enable_t2_unlimited":           &hcldec.AttrSpec{Name: "enable_t2_unlimited", Type: cty.Bool, Required: false},
		"iam_instance_profile":          &hcldec.AttrSpec{Name: "iam_instance_profile", Type: cty.String, Required: false},
		"skip_profile_validation":       &hcldec.AttrSpec{Name: "skip_profile_validation", Type: cty.Bool, Required: false},
		"temporary_iam_instance_profile_policy_document": &hcldec.BlockSpec{TypeName: "temporary_iam_instance_profile_policy_document", Nested: hcldec.ObjectSpec((*common.FlatPolicyDocument)(nil).HCL2Spec())},
		"shutdown_behavior":                     &hcldec.AttrSpec{Name: "shutdown_behavior", Type: cty.String, Required: false},
		"instance_type":                         &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
		"security_group_filter":                 &hcldec.BlockSpec{TypeName: "security_group_filter", Nested: hcldec.ObjectSpec((*common.FlatSecurityGroupFilterOptions)(nil).HCL2Spec())},
		"run_tags":                              &hcldec.AttrSpec{Name: "run_tags", Type: cty.Map(cty.String), Required: false},
		"run_tag":                               &hcldec.BlockListSpec{TypeName: "run_tag", Nested: hcldec.ObjectSpec((*config.FlatKeyValue)(nil).HCL2Spec())},
		"security_group_id":                     &hcldec.AttrSpec{Name: "security_group_id", Type: cty.String, Required: false},
		"security_group_ids":                    &hcldec.AttrSpec{Name: "security_group_ids", Type: cty.List(cty.String), Required: false},
		"source_ami":                            &hcldec.AttrSpec{Name: "source_ami", Type: cty.String, Required: false},
		"source_ami_filter":                     &hcldec.BlockSpec{TypeName: "source_ami_filter", Nested: hcldec.ObjectSpec((*common.FlatAmiFilterOptions)(nil).HCL2Spec())},
		"spot_instance_types":                   &hcldec.AttrSpec{Name: "spot_instance_types", Type: cty.List(cty.String), Required: false},
		"spot_price":                            &hcldec.AttrSpec{Name: "spot_price", Type: cty.String, Required: false},
		"spot_price_auto_product":               &hcldec.AttrSpec{Name: "spot_price_auto_product", Type: cty.String, Required: false},
		"spot_tags":                             &hcldec.AttrSpec{Name: "spot_tags", Type: cty.Map(cty.String), Required: false},
		"spot_tag":                              &hcldec.BlockListSpec{TypeName: "spot_tag", Nested: hcldec.ObjectSpec((*config.FlatKeyValue)(nil).HCL2Spec())},
		"subnet_filter":                         &hcldec.BlockSpec{TypeName: "subnet_filter", Nested: hcldec.ObjectSpec((*common.FlatSubnetFilterOptions)(nil).HCL2Spec())},
		"subnet_id":                             &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"tenancy":                               &hcldec.AttrSpec{Name: "tenancy", Type: cty.String, Required: false},
		"temporary_security_group_source_cidrs": &hcldec.AttrSpec{Name: "temporary_security_group_source_cidrs", Type: cty.List(cty.String), Required: false},
		"user_data":                             &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":                        &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"vpc_filter":                            &hcldec.BlockSpec{TypeName: "vpc_filter", Nested: hcldec.ObjectSpec((*common.FlatVpcFilterOptions)(nil).HCL2Spec())},
		"vpc_id":                                &hcldec.AttrSpec{Name: "vpc_id", Type: cty.String, Required: false},
		"windows_password_timeout":              &hcldec.AttrSpec{Name: "windows_password_timeout", Type: cty.String, Required: false},
		"communicator":                          &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":               &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                              &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                              &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                          &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                          &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":                      &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":               &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":               &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":               &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                           &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":             &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":           &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":                  &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":                  &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                               &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                           &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":                      &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":                        &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding":          &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":                &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":                      &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":                      &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":                &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":                  &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":                  &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":               &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file":          &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file":          &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":              &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":                        &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":                        &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":                    &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":                    &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":               &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":                &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":                    &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":                     &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":                        &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":                       &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":                        &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":                        &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                            &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":                        &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                            &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                         &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                         &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":                        &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":                        &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"ssh_interface":                         &hcldec.AttrSpec{Name: "ssh_interface", Type: cty.String, Required: false},
		"pause_before_ssm":                      &hcldec.AttrSpec{Name: "pause_before_ssm", Type: cty.String, Required: false},
		"session_manager_port":                  &hcldec.AttrSpec{Name: "session_manager_port", Type: cty.Number, Required: false},
		"ena_support":                           &hcldec.AttrSpec{Name: "ena_support", Type: cty.Bool, Required: false},
		"sriov_support":                         &hcldec.AttrSpec{Name: "sriov_support", Type: cty.Bool, Required: false},
		"ebs_volumes":                           &hcldec.BlockListSpec{TypeName: "ebs_volumes", Nested: hcldec.ObjectSpec((*FlatBlockDevice)(nil).HCL2Spec())},
		"run_volume_tags":                       &hcldec.AttrSpec{Name: "run_volume_tags", Type: cty.Map(cty.String), Required: false},
		"run_volume_tag":                        &hcldec.BlockListSpec{TypeName: "run_volume_tag", Nested: hcldec.ObjectSpec((*config.FlatKeyValue)(nil).HCL2Spec())},
	}
	return s
}
