/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// LayerDirective : - DIRECTIVE_UNSPECIFIED: Default value for unsupported/missing directive  - MAINTAINER: https://docs.docker.com/reference/builder/#maintainer  - RUN: https://docs.docker.com/reference/builder/#run  - CMD: https://docs.docker.com/reference/builder/#cmd  - LABEL: https://docs.docker.com/reference/builder/#label  - EXPOSE: https://docs.docker.com/reference/builder/#expose  - ENV: https://docs.docker.com/reference/builder/#env  - ADD: https://docs.docker.com/reference/builder/#add  - COPY: https://docs.docker.com/reference/builder/#copy  - ENTRYPOINT: https://docs.docker.com/reference/builder/#entrypoint  - VOLUME: https://docs.docker.com/reference/builder/#volume  - USER: https://docs.docker.com/reference/builder/#user  - WORKDIR: https://docs.docker.com/reference/builder/#workdir  - ARG: https://docs.docker.com/reference/builder/#arg  - ONBUILD: https://docs.docker.com/reference/builder/#onbuild  - STOPSIGNAL: https://docs.docker.com/reference/builder/#stopsignal  - HEALTHCHECK: https://docs.docker.com/reference/builder/#healthcheck  - SHELL: https://docs.docker.com/reference/builder/#shell
type LayerDirective string

// List of LayerDirective
const (
	DIRECTIVE_UNSPECIFIED LayerDirective = "DIRECTIVE_UNSPECIFIED"
	MAINTAINER LayerDirective = "MAINTAINER"
	RUN LayerDirective = "RUN"
	CMD LayerDirective = "CMD"
	LABEL LayerDirective = "LABEL"
	EXPOSE LayerDirective = "EXPOSE"
	ENV LayerDirective = "ENV"
	ADD LayerDirective = "ADD"
	COPY LayerDirective = "COPY"
	ENTRYPOINT LayerDirective = "ENTRYPOINT"
	VOLUME LayerDirective = "VOLUME"
	USER LayerDirective = "USER"
	WORKDIR LayerDirective = "WORKDIR"
	ARG LayerDirective = "ARG"
	ONBUILD LayerDirective = "ONBUILD"
	STOPSIGNAL LayerDirective = "STOPSIGNAL"
	HEALTHCHECK LayerDirective = "HEALTHCHECK"
	SHELL LayerDirective = "SHELL"
)
