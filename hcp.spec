# Copyright 2024 Wong Hoi Sing Edison <hswong3i@pantarei-design.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

%global debug_package %{nil}

%global source_date_epoch_from_changelog 0

Name: hcp
Epoch: 100
Version: 0.7.0
Release: 1%{?dist}
Summary: HCP Command-Line Interface
License: MPL-2.0
URL: https://github.com/hashicorp/hcp/tags
Source0: %{name}_%{version}.orig.tar.gz
BuildRequires: golang-1.23
BuildRequires: glibc-static

%description
The hcp CLI lets you administer HashiCorp Cloud Platform (HCP) resources
and services. You can interact with hcp directly or integrate it in
scripts to automate your workflows.

%prep
%autosetup -T -c -n %{name}_%{version}-%{release}
tar -zx -f %{S:0} --strip-components=1 -C .

%build
mkdir -p bin
set -ex && \
    export CGO_ENABLED=0 && \
    go build \
        -mod vendor -buildmode pie -v \
        -ldflags "-s -w -extldflags '-static -lm'" \
        -o ./bin/hcp .

%install
install -Dpm755 -d %{buildroot}%{_bindir}
install -Dpm755 -d %{buildroot}%{_prefix}/share/bash-completion/completions
install -Dpm755 -t %{buildroot}%{_bindir}/ bin/hcp
install -Dpm644 -t %{buildroot}%{_prefix}/share/bash-completion/completions contrib/bash-completion/hcp

%files
%license LICENSE
%{_bindir}/*
%{_prefix}/share/bash-completion/completions/*

%changelog
