package tpl

// Tpl3 ...
const Tpl3 = `
{
	"name": "core",
	"version": "0.0.0",
	"lockfileVersion": 1,
	"requires": true,
	"dependencies": {
	  "@babel/cli": {
		"version": "7.7.7",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fcli/-/cli-7.7.7.tgz",
		"integrity": "sha512-XQw5KyCZyu/M8/0rYiZyuwbgIQNzOrJzs9dDLX+MieSgBwTLvTj4QVbLmxJACAIvQIDT7PtyHN2sC48EOWTgaA==",
		"dev": true,
		"requires": {
		  "chokidar": "^2.1.8",
		  "commander": "^4.0.1",
		  "convert-source-map": "^1.1.0",
		  "fs-readdir-recursive": "^1.1.0",
		  "glob": "^7.0.0",
		  "lodash": "^4.17.13",
		  "make-dir": "^2.1.0",
		  "slash": "^2.0.0",
		  "source-map": "^0.5.0"
		}
	  },
	  "@babel/code-frame": {
		"version": "7.5.5",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fcode-frame/-/code-frame-7.5.5.tgz",
		"integrity": "sha512-27d4lZoomVyo51VegxI20xZPuSHusqbQag/ztrBC7wegWoQ1nLREPVSKSW8byhTlzTKyNE4ifaTA6lCp7JjpFw==",
		"dev": true,
		"requires": {
		  "@babel/highlight": "^7.0.0"
		}
	  },
	  "@babel/core": {
		"version": "7.7.7",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fcore/-/core-7.7.7.tgz",
		"integrity": "sha512-jlSjuj/7z138NLZALxVgrx13AOtqip42ATZP7+kYl53GvDV6+4dCek1mVUo8z8c8Xnw/mx2q3d9HWh3griuesQ==",
		"dev": true,
		"requires": {
		  "@babel/code-frame": "^7.5.5",
		  "@babel/generator": "^7.7.7",
		  "@babel/helpers": "^7.7.4",
		  "@babel/parser": "^7.7.7",
		  "@babel/template": "^7.7.4",
		  "@babel/traverse": "^7.7.4",
		  "@babel/types": "^7.7.4",
		  "convert-source-map": "^1.7.0",
		  "debug": "^4.1.0",
		  "json5": "^2.1.0",
		  "lodash": "^4.17.13",
		  "resolve": "^1.3.2",
		  "semver": "^5.4.1",
		  "source-map": "^0.5.0"
		}
	  },
	  "@babel/generator": {
		"version": "7.7.7",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fgenerator/-/generator-7.7.7.tgz",
		"integrity": "sha512-/AOIBpHh/JU1l0ZFS4kiRCBnLi6OTHzh0RPk3h9isBxkkqELtQNFi1Vr/tiG9p1yfoUdKVwISuXWQR+hwwM4VQ==",
		"dev": true,
		"requires": {
		  "@babel/types": "^7.7.4",
		  "jsesc": "^2.5.1",
		  "lodash": "^4.17.13",
		  "source-map": "^0.5.0"
		}
	  },
	  "@babel/helper-annotate-as-pure": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-annotate-as-pure/-/helper-annotate-as-pure-7.7.4.tgz",
		"integrity": "sha512-2BQmQgECKzYKFPpiycoF9tlb5HA4lrVyAmLLVK177EcQAqjVLciUb2/R+n1boQ9y5ENV3uz2ZqiNw7QMBBw1Og==",
		"dev": true,
		"requires": {
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-builder-binary-assignment-operator-visitor": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-builder-binary-assignment-operator-visitor/-/helper-builder-binary-assignment-operator-visitor-7.7.4.tgz",
		"integrity": "sha512-Biq/d/WtvfftWZ9Uf39hbPBYDUo986m5Bb4zhkeYDGUllF43D+nUe5M6Vuo6/8JDK/0YX/uBdeoQpyaNhNugZQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-explode-assignable-expression": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-call-delegate": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-call-delegate/-/helper-call-delegate-7.7.4.tgz",
		"integrity": "sha512-8JH9/B7J7tCYJ2PpWVpw9JhPuEVHztagNVuQAFBVFYluRMlpG7F1CgKEgGeL6KFqcsIa92ZYVj6DSc0XwmN1ZA==",
		"dev": true,
		"requires": {
		  "@babel/helper-hoist-variables": "^7.7.4",
		  "@babel/traverse": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-create-class-features-plugin": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-create-class-features-plugin/-/helper-create-class-features-plugin-7.7.4.tgz",
		"integrity": "sha512-l+OnKACG4uiDHQ/aJT8dwpR+LhCJALxL0mJ6nzjB25e5IPwqV1VOsY7ah6UB1DG+VOXAIMtuC54rFJGiHkxjgA==",
		"dev": true,
		"requires": {
		  "@babel/helper-function-name": "^7.7.4",
		  "@babel/helper-member-expression-to-functions": "^7.7.4",
		  "@babel/helper-optimise-call-expression": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/helper-replace-supers": "^7.7.4",
		  "@babel/helper-split-export-declaration": "^7.7.4"
		}
	  },
	  "@babel/helper-create-regexp-features-plugin": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-create-regexp-features-plugin/-/helper-create-regexp-features-plugin-7.7.4.tgz",
		"integrity": "sha512-Mt+jBKaxL0zfOIWrfQpnfYCN7/rS6GKx6CCCfuoqVVd+17R8zNDlzVYmIi9qyb2wOk002NsmSTDymkIygDUH7A==",
		"dev": true,
		"requires": {
		  "@babel/helper-regex": "^7.4.4",
		  "regexpu-core": "^4.6.0"
		}
	  },
	  "@babel/helper-define-map": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-define-map/-/helper-define-map-7.7.4.tgz",
		"integrity": "sha512-v5LorqOa0nVQUvAUTUF3KPastvUt/HzByXNamKQ6RdJRTV7j8rLL+WB5C/MzzWAwOomxDhYFb1wLLxHqox86lg==",
		"dev": true,
		"requires": {
		  "@babel/helper-function-name": "^7.7.4",
		  "@babel/types": "^7.7.4",
		  "lodash": "^4.17.13"
		}
	  },
	  "@babel/helper-explode-assignable-expression": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-explode-assignable-expression/-/helper-explode-assignable-expression-7.7.4.tgz",
		"integrity": "sha512-2/SicuFrNSXsZNBxe5UGdLr+HZg+raWBLE9vC98bdYOKX/U6PY0mdGlYUJdtTDPSU0Lw0PNbKKDpwYHJLn2jLg==",
		"dev": true,
		"requires": {
		  "@babel/traverse": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-function-name": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-function-name/-/helper-function-name-7.7.4.tgz",
		"integrity": "sha512-AnkGIdiBhEuiwdoMnKm7jfPfqItZhgRaZfMg1XX3bS25INOnLPjPG1Ppnajh8eqgt5kPJnfqrRHqFqmjKDZLzQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-get-function-arity": "^7.7.4",
		  "@babel/template": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-get-function-arity": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-get-function-arity/-/helper-get-function-arity-7.7.4.tgz",
		"integrity": "sha512-QTGKEdCkjgzgfJ3bAyRwF4yyT3pg+vDgan8DSivq1eS0gwi+KGKE5x8kRcbeFTb/673mkO5SN1IZfmCfA5o+EA==",
		"dev": true,
		"requires": {
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-hoist-variables": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-hoist-variables/-/helper-hoist-variables-7.7.4.tgz",
		"integrity": "sha512-wQC4xyvc1Jo/FnLirL6CEgPgPCa8M74tOdjWpRhQYapz5JC7u3NYU1zCVoVAGCE3EaIP9T1A3iW0WLJ+reZlpQ==",
		"dev": true,
		"requires": {
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-member-expression-to-functions": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-member-expression-to-functions/-/helper-member-expression-to-functions-7.7.4.tgz",
		"integrity": "sha512-9KcA1X2E3OjXl/ykfMMInBK+uVdfIVakVe7W7Lg3wfXUNyS3Q1HWLFRwZIjhqiCGbslummPDnmb7vIekS0C1vw==",
		"dev": true,
		"requires": {
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-module-imports": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-module-imports/-/helper-module-imports-7.7.4.tgz",
		"integrity": "sha512-dGcrX6K9l8258WFjyDLJwuVKxR4XZfU0/vTUgOQYWEnRD8mgr+p4d6fCUMq/ys0h4CCt/S5JhbvtyErjWouAUQ==",
		"dev": true,
		"requires": {
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-module-transforms": {
		"version": "7.7.5",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-module-transforms/-/helper-module-transforms-7.7.5.tgz",
		"integrity": "sha512-A7pSxyJf1gN5qXVcidwLWydjftUN878VkalhXX5iQDuGyiGK3sOrrKKHF4/A4fwHtnsotv/NipwAeLzY4KQPvw==",
		"dev": true,
		"requires": {
		  "@babel/helper-module-imports": "^7.7.4",
		  "@babel/helper-simple-access": "^7.7.4",
		  "@babel/helper-split-export-declaration": "^7.7.4",
		  "@babel/template": "^7.7.4",
		  "@babel/types": "^7.7.4",
		  "lodash": "^4.17.13"
		}
	  },
	  "@babel/helper-optimise-call-expression": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-optimise-call-expression/-/helper-optimise-call-expression-7.7.4.tgz",
		"integrity": "sha512-VB7gWZ2fDkSuqW6b1AKXkJWO5NyNI3bFL/kK79/30moK57blr6NbH8xcl2XcKCwOmJosftWunZqfO84IGq3ZZg==",
		"dev": true,
		"requires": {
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-plugin-utils": {
		"version": "7.0.0",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-plugin-utils/-/helper-plugin-utils-7.0.0.tgz",
		"integrity": "sha512-CYAOUCARwExnEixLdB6sDm2dIJ/YgEAKDM1MOeMeZu9Ld/bDgVo8aiWrXwcY7OBh+1Ea2uUcVRcxKk0GJvW7QA==",
		"dev": true
	  },
	  "@babel/helper-regex": {
		"version": "7.5.5",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-regex/-/helper-regex-7.5.5.tgz",
		"integrity": "sha512-CkCYQLkfkiugbRDO8eZn6lRuR8kzZoGXCg3149iTk5se7g6qykSpy3+hELSwquhu+TgHn8nkLiBwHvNX8Hofcw==",
		"dev": true,
		"requires": {
		  "lodash": "^4.17.13"
		}
	  },
	  "@babel/helper-remap-async-to-generator": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-remap-async-to-generator/-/helper-remap-async-to-generator-7.7.4.tgz",
		"integrity": "sha512-Sk4xmtVdM9sA/jCI80f+KS+Md+ZHIpjuqmYPk1M7F/upHou5e4ReYmExAiu6PVe65BhJPZA2CY9x9k4BqE5klw==",
		"dev": true,
		"requires": {
		  "@babel/helper-annotate-as-pure": "^7.7.4",
		  "@babel/helper-wrap-function": "^7.7.4",
		  "@babel/template": "^7.7.4",
		  "@babel/traverse": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-replace-supers": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-replace-supers/-/helper-replace-supers-7.7.4.tgz",
		"integrity": "sha512-pP0tfgg9hsZWo5ZboYGuBn/bbYT/hdLPVSS4NMmiRJdwWhP0IznPwN9AE1JwyGsjSPLC364I0Qh5p+EPkGPNpg==",
		"dev": true,
		"requires": {
		  "@babel/helper-member-expression-to-functions": "^7.7.4",
		  "@babel/helper-optimise-call-expression": "^7.7.4",
		  "@babel/traverse": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-simple-access": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-simple-access/-/helper-simple-access-7.7.4.tgz",
		"integrity": "sha512-zK7THeEXfan7UlWsG2A6CI/L9jVnI5+xxKZOdej39Y0YtDYKx9raHk5F2EtK9K8DHRTihYwg20ADt9S36GR78A==",
		"dev": true,
		"requires": {
		  "@babel/template": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-split-export-declaration": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-split-export-declaration/-/helper-split-export-declaration-7.7.4.tgz",
		"integrity": "sha512-guAg1SXFcVr04Guk9eq0S4/rWS++sbmyqosJzVs8+1fH5NI+ZcmkaSkc7dmtAFbHFva6yRJnjW3yAcGxjueDug==",
		"dev": true,
		"requires": {
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helper-wrap-function": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelper-wrap-function/-/helper-wrap-function-7.7.4.tgz",
		"integrity": "sha512-VsfzZt6wmsocOaVU0OokwrIytHND55yvyT4BPB9AIIgwr8+x7617hetdJTsuGwygN5RC6mxA9EJztTjuwm2ofg==",
		"dev": true,
		"requires": {
		  "@babel/helper-function-name": "^7.7.4",
		  "@babel/template": "^7.7.4",
		  "@babel/traverse": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/helpers": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhelpers/-/helpers-7.7.4.tgz",
		"integrity": "sha512-ak5NGZGJ6LV85Q1Zc9gn2n+ayXOizryhjSUBTdu5ih1tlVCJeuQENzc4ItyCVhINVXvIT/ZQ4mheGIsfBkpskg==",
		"dev": true,
		"requires": {
		  "@babel/template": "^7.7.4",
		  "@babel/traverse": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/highlight": {
		"version": "7.5.0",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fhighlight/-/highlight-7.5.0.tgz",
		"integrity": "sha512-7dV4eu9gBxoM0dAnj/BCFDW9LFU0zvTrkq0ugM7pnHEgguOEeOz1so2ZghEdzviYzQEED0r4EAgpsBChKy1TRQ==",
		"dev": true,
		"requires": {
		  "chalk": "^2.0.0",
		  "esutils": "^2.0.2",
		  "js-tokens": "^4.0.0"
		}
	  },
	  "@babel/parser": {
		"version": "7.7.7",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fparser/-/parser-7.7.7.tgz",
		"integrity": "sha512-WtTZMZAZLbeymhkd/sEaPD8IQyGAhmuTuvTzLiCFM7iXiVdY0gc0IaI+cW0fh1BnSMbJSzXX6/fHllgHKwHhXw==",
		"dev": true
	  },
	  "@babel/plugin-proposal-async-generator-functions": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-proposal-async-generator-functions/-/plugin-proposal-async-generator-functions-7.7.4.tgz",
		"integrity": "sha512-1ypyZvGRXriY/QP668+s8sFr2mqinhkRDMPSQLNghCQE+GAkFtp+wkHVvg2+Hdki8gwP+NFzJBJ/N1BfzCCDEw==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/helper-remap-async-to-generator": "^7.7.4",
		  "@babel/plugin-syntax-async-generators": "^7.7.4"
		}
	  },
	  "@babel/plugin-proposal-class-properties": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-proposal-class-properties/-/plugin-proposal-class-properties-7.7.4.tgz",
		"integrity": "sha512-EcuXeV4Hv1X3+Q1TsuOmyyxeTRiSqurGJ26+I/FW1WbymmRRapVORm6x1Zl3iDIHyRxEs+VXWp6qnlcfcJSbbw==",
		"dev": true,
		"requires": {
		  "@babel/helper-create-class-features-plugin": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-proposal-dynamic-import": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-proposal-dynamic-import/-/plugin-proposal-dynamic-import-7.7.4.tgz",
		"integrity": "sha512-StH+nGAdO6qDB1l8sZ5UBV8AC3F2VW2I8Vfld73TMKyptMU9DY5YsJAS8U81+vEtxcH3Y/La0wG0btDrhpnhjQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/plugin-syntax-dynamic-import": "^7.7.4"
		}
	  },
	  "@babel/plugin-proposal-json-strings": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-proposal-json-strings/-/plugin-proposal-json-strings-7.7.4.tgz",
		"integrity": "sha512-wQvt3akcBTfLU/wYoqm/ws7YOAQKu8EVJEvHip/mzkNtjaclQoCCIqKXFP5/eyfnfbQCDV3OLRIK3mIVyXuZlw==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/plugin-syntax-json-strings": "^7.7.4"
		}
	  },
	  "@babel/plugin-proposal-object-rest-spread": {
		"version": "7.7.7",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-proposal-object-rest-spread/-/plugin-proposal-object-rest-spread-7.7.7.tgz",
		"integrity": "sha512-3qp9I8lelgzNedI3hrhkvhaEYree6+WHnyA/q4Dza9z7iEIs1eyhWyJnetk3jJ69RT0AT4G0UhEGwyGFJ7GUuQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/plugin-syntax-object-rest-spread": "^7.7.4"
		}
	  },
	  "@babel/plugin-proposal-optional-catch-binding": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-proposal-optional-catch-binding/-/plugin-proposal-optional-catch-binding-7.7.4.tgz",
		"integrity": "sha512-DyM7U2bnsQerCQ+sejcTNZh8KQEUuC3ufzdnVnSiUv/qoGJp2Z3hanKL18KDhsBT5Wj6a7CMT5mdyCNJsEaA9w==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/plugin-syntax-optional-catch-binding": "^7.7.4"
		}
	  },
	  "@babel/plugin-proposal-unicode-property-regex": {
		"version": "7.7.7",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-proposal-unicode-property-regex/-/plugin-proposal-unicode-property-regex-7.7.7.tgz",
		"integrity": "sha512-80PbkKyORBUVm1fbTLrHpYdJxMThzM1UqFGh0ALEhO9TYbG86Ah9zQYAB/84axz2vcxefDLdZwWwZNlYARlu9w==",
		"dev": true,
		"requires": {
		  "@babel/helper-create-regexp-features-plugin": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-syntax-async-generators": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-syntax-async-generators/-/plugin-syntax-async-generators-7.7.4.tgz",
		"integrity": "sha512-Li4+EjSpBgxcsmeEF8IFcfV/+yJGxHXDirDkEoyFjumuwbmfCVHUt0HuowD/iGM7OhIRyXJH9YXxqiH6N815+g==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-syntax-dynamic-import": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-syntax-dynamic-import/-/plugin-syntax-dynamic-import-7.7.4.tgz",
		"integrity": "sha512-jHQW0vbRGvwQNgyVxwDh4yuXu4bH1f5/EICJLAhl1SblLs2CDhrsmCk+v5XLdE9wxtAFRyxx+P//Iw+a5L/tTg==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-syntax-json-strings": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-syntax-json-strings/-/plugin-syntax-json-strings-7.7.4.tgz",
		"integrity": "sha512-QpGupahTQW1mHRXddMG5srgpHWqRLwJnJZKXTigB9RPFCCGbDGCgBeM/iC82ICXp414WeYx/tD54w7M2qRqTMg==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-syntax-object-rest-spread": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-syntax-object-rest-spread/-/plugin-syntax-object-rest-spread-7.7.4.tgz",
		"integrity": "sha512-mObR+r+KZq0XhRVS2BrBKBpr5jqrqzlPvS9C9vuOf5ilSwzloAl7RPWLrgKdWS6IreaVrjHxTjtyqFiOisaCwg==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-syntax-optional-catch-binding": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-syntax-optional-catch-binding/-/plugin-syntax-optional-catch-binding-7.7.4.tgz",
		"integrity": "sha512-4ZSuzWgFxqHRE31Glu+fEr/MirNZOMYmD/0BhBWyLyOOQz/gTAl7QmWm2hX1QxEIXsr2vkdlwxIzTyiYRC4xcQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-syntax-top-level-await": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-syntax-top-level-await/-/plugin-syntax-top-level-await-7.7.4.tgz",
		"integrity": "sha512-wdsOw0MvkL1UIgiQ/IFr3ETcfv1xb8RMM0H9wbiDyLaJFyiDg5oZvDLCXosIXmFeIlweML5iOBXAkqddkYNizg==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-arrow-functions": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-arrow-functions/-/plugin-transform-arrow-functions-7.7.4.tgz",
		"integrity": "sha512-zUXy3e8jBNPiffmqkHRNDdZM2r8DWhCB7HhcoyZjiK1TxYEluLHAvQuYnTT+ARqRpabWqy/NHkO6e3MsYB5YfA==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-async-to-generator": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-async-to-generator/-/plugin-transform-async-to-generator-7.7.4.tgz",
		"integrity": "sha512-zpUTZphp5nHokuy8yLlyafxCJ0rSlFoSHypTUWgpdwoDXWQcseaect7cJ8Ppk6nunOM6+5rPMkod4OYKPR5MUg==",
		"dev": true,
		"requires": {
		  "@babel/helper-module-imports": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/helper-remap-async-to-generator": "^7.7.4"
		}
	  },
	  "@babel/plugin-transform-block-scoped-functions": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-block-scoped-functions/-/plugin-transform-block-scoped-functions-7.7.4.tgz",
		"integrity": "sha512-kqtQzwtKcpPclHYjLK//3lH8OFsCDuDJBaFhVwf8kqdnF6MN4l618UDlcA7TfRs3FayrHj+svYnSX8MC9zmUyQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-block-scoping": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-block-scoping/-/plugin-transform-block-scoping-7.7.4.tgz",
		"integrity": "sha512-2VBe9u0G+fDt9B5OV5DQH4KBf5DoiNkwFKOz0TCvBWvdAN2rOykCTkrL+jTLxfCAm76l9Qo5OqL7HBOx2dWggg==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "lodash": "^4.17.13"
		}
	  },
	  "@babel/plugin-transform-classes": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-classes/-/plugin-transform-classes-7.7.4.tgz",
		"integrity": "sha512-sK1mjWat7K+buWRuImEzjNf68qrKcrddtpQo3swi9j7dUcG6y6R6+Di039QN2bD1dykeswlagupEmpOatFHHUg==",
		"dev": true,
		"requires": {
		  "@babel/helper-annotate-as-pure": "^7.7.4",
		  "@babel/helper-define-map": "^7.7.4",
		  "@babel/helper-function-name": "^7.7.4",
		  "@babel/helper-optimise-call-expression": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/helper-replace-supers": "^7.7.4",
		  "@babel/helper-split-export-declaration": "^7.7.4",
		  "globals": "^11.1.0"
		}
	  },
	  "@babel/plugin-transform-computed-properties": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-computed-properties/-/plugin-transform-computed-properties-7.7.4.tgz",
		"integrity": "sha512-bSNsOsZnlpLLyQew35rl4Fma3yKWqK3ImWMSC/Nc+6nGjC9s5NFWAer1YQ899/6s9HxO2zQC1WoFNfkOqRkqRQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-destructuring": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-destructuring/-/plugin-transform-destructuring-7.7.4.tgz",
		"integrity": "sha512-4jFMXI1Cu2aXbcXXl8Lr6YubCn6Oc7k9lLsu8v61TZh+1jny2BWmdtvY9zSUlLdGUvcy9DMAWyZEOqjsbeg/wA==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-dotall-regex": {
		"version": "7.7.7",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-dotall-regex/-/plugin-transform-dotall-regex-7.7.7.tgz",
		"integrity": "sha512-b4in+YlTeE/QmTgrllnb3bHA0HntYvjz8O3Mcbx75UBPJA2xhb5A8nle498VhxSXJHQefjtQxpnLPehDJ4TRlg==",
		"dev": true,
		"requires": {
		  "@babel/helper-create-regexp-features-plugin": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-duplicate-keys": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-duplicate-keys/-/plugin-transform-duplicate-keys-7.7.4.tgz",
		"integrity": "sha512-g1y4/G6xGWMD85Tlft5XedGaZBCIVN+/P0bs6eabmcPP9egFleMAo65OOjlhcz1njpwagyY3t0nsQC9oTFegJA==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-exponentiation-operator": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-exponentiation-operator/-/plugin-transform-exponentiation-operator-7.7.4.tgz",
		"integrity": "sha512-MCqiLfCKm6KEA1dglf6Uqq1ElDIZwFuzz1WH5mTf8k2uQSxEJMbOIEh7IZv7uichr7PMfi5YVSrr1vz+ipp7AQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-builder-binary-assignment-operator-visitor": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-for-of": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-for-of/-/plugin-transform-for-of-7.7.4.tgz",
		"integrity": "sha512-zZ1fD1B8keYtEcKF+M1TROfeHTKnijcVQm0yO/Yu1f7qoDoxEIc/+GX6Go430Bg84eM/xwPFp0+h4EbZg7epAA==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-function-name": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-function-name/-/plugin-transform-function-name-7.7.4.tgz",
		"integrity": "sha512-E/x09TvjHNhsULs2IusN+aJNRV5zKwxu1cpirZyRPw+FyyIKEHPXTsadj48bVpc1R5Qq1B5ZkzumuFLytnbT6g==",
		"dev": true,
		"requires": {
		  "@babel/helper-function-name": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-literals": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-literals/-/plugin-transform-literals-7.7.4.tgz",
		"integrity": "sha512-X2MSV7LfJFm4aZfxd0yLVFrEXAgPqYoDG53Br/tCKiKYfX0MjVjQeWPIhPHHsCqzwQANq+FLN786fF5rgLS+gw==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-member-expression-literals": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-member-expression-literals/-/plugin-transform-member-expression-literals-7.7.4.tgz",
		"integrity": "sha512-9VMwMO7i69LHTesL0RdGy93JU6a+qOPuvB4F4d0kR0zyVjJRVJRaoaGjhtki6SzQUu8yen/vxPKN6CWnCUw6bA==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-modules-amd": {
		"version": "7.7.5",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-modules-amd/-/plugin-transform-modules-amd-7.7.5.tgz",
		"integrity": "sha512-CT57FG4A2ZUNU1v+HdvDSDrjNWBrtCmSH6YbbgN3Lrf0Di/q/lWRxZrE72p3+HCCz9UjfZOEBdphgC0nzOS6DQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-module-transforms": "^7.7.5",
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "babel-plugin-dynamic-import-node": "^2.3.0"
		}
	  },
	  "@babel/plugin-transform-modules-commonjs": {
		"version": "7.7.5",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-modules-commonjs/-/plugin-transform-modules-commonjs-7.7.5.tgz",
		"integrity": "sha512-9Cq4zTFExwFhQI6MT1aFxgqhIsMWQWDVwOgLzl7PTWJHsNaqFvklAU+Oz6AQLAS0dJKTwZSOCo20INwktxpi3Q==",
		"dev": true,
		"requires": {
		  "@babel/helper-module-transforms": "^7.7.5",
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/helper-simple-access": "^7.7.4",
		  "babel-plugin-dynamic-import-node": "^2.3.0"
		}
	  },
	  "@babel/plugin-transform-modules-systemjs": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-modules-systemjs/-/plugin-transform-modules-systemjs-7.7.4.tgz",
		"integrity": "sha512-y2c96hmcsUi6LrMqvmNDPBBiGCiQu0aYqpHatVVu6kD4mFEXKjyNxd/drc18XXAf9dv7UXjrZwBVmTTGaGP8iw==",
		"dev": true,
		"requires": {
		  "@babel/helper-hoist-variables": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "babel-plugin-dynamic-import-node": "^2.3.0"
		}
	  },
	  "@babel/plugin-transform-modules-umd": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-modules-umd/-/plugin-transform-modules-umd-7.7.4.tgz",
		"integrity": "sha512-u2B8TIi0qZI4j8q4C51ktfO7E3cQ0qnaXFI1/OXITordD40tt17g/sXqgNNCcMTcBFKrUPcGDx+TBJuZxLx7tw==",
		"dev": true,
		"requires": {
		  "@babel/helper-module-transforms": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-named-capturing-groups-regex": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-named-capturing-groups-regex/-/plugin-transform-named-capturing-groups-regex-7.7.4.tgz",
		"integrity": "sha512-jBUkiqLKvUWpv9GLSuHUFYdmHg0ujC1JEYoZUfeOOfNydZXp1sXObgyPatpcwjWgsdBGsagWW0cdJpX/DO2jMw==",
		"dev": true,
		"requires": {
		  "@babel/helper-create-regexp-features-plugin": "^7.7.4"
		}
	  },
	  "@babel/plugin-transform-new-target": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-new-target/-/plugin-transform-new-target-7.7.4.tgz",
		"integrity": "sha512-CnPRiNtOG1vRodnsyGX37bHQleHE14B9dnnlgSeEs3ek3fHN1A1SScglTCg1sfbe7sRQ2BUcpgpTpWSfMKz3gg==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-object-super": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-object-super/-/plugin-transform-object-super-7.7.4.tgz",
		"integrity": "sha512-ho+dAEhC2aRnff2JCA0SAK7V2R62zJd/7dmtoe7MHcso4C2mS+vZjn1Pb1pCVZvJs1mgsvv5+7sT+m3Bysb6eg==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/helper-replace-supers": "^7.7.4"
		}
	  },
	  "@babel/plugin-transform-parameters": {
		"version": "7.7.7",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-parameters/-/plugin-transform-parameters-7.7.7.tgz",
		"integrity": "sha512-OhGSrf9ZBrr1fw84oFXj5hgi8Nmg+E2w5L7NhnG0lPvpDtqd7dbyilM2/vR8CKbJ907RyxPh2kj6sBCSSfI9Ew==",
		"dev": true,
		"requires": {
		  "@babel/helper-call-delegate": "^7.7.4",
		  "@babel/helper-get-function-arity": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-property-literals": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-property-literals/-/plugin-transform-property-literals-7.7.4.tgz",
		"integrity": "sha512-MatJhlC4iHsIskWYyawl53KuHrt+kALSADLQQ/HkhTjX954fkxIEh4q5slL4oRAnsm/eDoZ4q0CIZpcqBuxhJQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-regenerator": {
		"version": "7.7.5",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-regenerator/-/plugin-transform-regenerator-7.7.5.tgz",
		"integrity": "sha512-/8I8tPvX2FkuEyWbjRCt4qTAgZK0DVy8QRguhA524UH48RfGJy94On2ri+dCuwOpcerPRl9O4ebQkRcVzIaGBw==",
		"dev": true,
		"requires": {
		  "regenerator-transform": "^0.14.0"
		}
	  },
	  "@babel/plugin-transform-reserved-words": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-reserved-words/-/plugin-transform-reserved-words-7.7.4.tgz",
		"integrity": "sha512-OrPiUB5s5XvkCO1lS7D8ZtHcswIC57j62acAnJZKqGGnHP+TIc/ljQSrgdX/QyOTdEK5COAhuc820Hi1q2UgLQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-runtime": {
		"version": "7.7.6",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-runtime/-/plugin-transform-runtime-7.7.6.tgz",
		"integrity": "sha512-tajQY+YmXR7JjTwRvwL4HePqoL3DYxpYXIHKVvrOIvJmeHe2y1w4tz5qz9ObUDC9m76rCzIMPyn4eERuwA4a4A==",
		"dev": true,
		"requires": {
		  "@babel/helper-module-imports": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "resolve": "^1.8.1",
		  "semver": "^5.5.1"
		}
	  },
	  "@babel/plugin-transform-shorthand-properties": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-shorthand-properties/-/plugin-transform-shorthand-properties-7.7.4.tgz",
		"integrity": "sha512-q+suddWRfIcnyG5YiDP58sT65AJDZSUhXQDZE3r04AuqD6d/XLaQPPXSBzP2zGerkgBivqtQm9XKGLuHqBID6Q==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-spread": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-spread/-/plugin-transform-spread-7.7.4.tgz",
		"integrity": "sha512-8OSs0FLe5/80cndziPlg4R0K6HcWSM0zyNhHhLsmw/Nc5MaA49cAsnoJ/t/YZf8qkG7fD+UjTRaApVDB526d7Q==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-sticky-regex": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-sticky-regex/-/plugin-transform-sticky-regex-7.7.4.tgz",
		"integrity": "sha512-Ls2NASyL6qtVe1H1hXts9yuEeONV2TJZmplLONkMPUG158CtmnrzW5Q5teibM5UVOFjG0D3IC5mzXR6pPpUY7A==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/helper-regex": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-template-literals": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-template-literals/-/plugin-transform-template-literals-7.7.4.tgz",
		"integrity": "sha512-sA+KxLwF3QwGj5abMHkHgshp9+rRz+oY9uoRil4CyLtgEuE/88dpkeWgNk5qKVsJE9iSfly3nvHapdRiIS2wnQ==",
		"dev": true,
		"requires": {
		  "@babel/helper-annotate-as-pure": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-typeof-symbol": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-typeof-symbol/-/plugin-transform-typeof-symbol-7.7.4.tgz",
		"integrity": "sha512-KQPUQ/7mqe2m0B8VecdyaW5XcQYaePyl9R7IsKd+irzj6jvbhoGnRE+M0aNkyAzI07VfUQ9266L5xMARitV3wg==",
		"dev": true,
		"requires": {
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/plugin-transform-unicode-regex": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fplugin-transform-unicode-regex/-/plugin-transform-unicode-regex-7.7.4.tgz",
		"integrity": "sha512-N77UUIV+WCvE+5yHw+oks3m18/umd7y392Zv7mYTpFqHtkpcc+QUz+gLJNTWVlWROIWeLqY0f3OjZxV5TcXnRw==",
		"dev": true,
		"requires": {
		  "@babel/helper-create-regexp-features-plugin": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0"
		}
	  },
	  "@babel/preset-env": {
		"version": "7.7.7",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2fpreset-env/-/preset-env-7.7.7.tgz",
		"integrity": "sha512-pCu0hrSSDVI7kCVUOdcMNQEbOPJ52E+LrQ14sN8uL2ALfSqePZQlKrOy+tM4uhEdYlCHi4imr8Zz2cZe9oSdIg==",
		"dev": true,
		"requires": {
		  "@babel/helper-module-imports": "^7.7.4",
		  "@babel/helper-plugin-utils": "^7.0.0",
		  "@babel/plugin-proposal-async-generator-functions": "^7.7.4",
		  "@babel/plugin-proposal-dynamic-import": "^7.7.4",
		  "@babel/plugin-proposal-json-strings": "^7.7.4",
		  "@babel/plugin-proposal-object-rest-spread": "^7.7.7",
		  "@babel/plugin-proposal-optional-catch-binding": "^7.7.4",
		  "@babel/plugin-proposal-unicode-property-regex": "^7.7.7",
		  "@babel/plugin-syntax-async-generators": "^7.7.4",
		  "@babel/plugin-syntax-dynamic-import": "^7.7.4",
		  "@babel/plugin-syntax-json-strings": "^7.7.4",
		  "@babel/plugin-syntax-object-rest-spread": "^7.7.4",
		  "@babel/plugin-syntax-optional-catch-binding": "^7.7.4",
		  "@babel/plugin-syntax-top-level-await": "^7.7.4",
		  "@babel/plugin-transform-arrow-functions": "^7.7.4",
		  "@babel/plugin-transform-async-to-generator": "^7.7.4",
		  "@babel/plugin-transform-block-scoped-functions": "^7.7.4",
		  "@babel/plugin-transform-block-scoping": "^7.7.4",
		  "@babel/plugin-transform-classes": "^7.7.4",
		  "@babel/plugin-transform-computed-properties": "^7.7.4",
		  "@babel/plugin-transform-destructuring": "^7.7.4",
		  "@babel/plugin-transform-dotall-regex": "^7.7.7",
		  "@babel/plugin-transform-duplicate-keys": "^7.7.4",
		  "@babel/plugin-transform-exponentiation-operator": "^7.7.4",
		  "@babel/plugin-transform-for-of": "^7.7.4",
		  "@babel/plugin-transform-function-name": "^7.7.4",
		  "@babel/plugin-transform-literals": "^7.7.4",
		  "@babel/plugin-transform-member-expression-literals": "^7.7.4",
		  "@babel/plugin-transform-modules-amd": "^7.7.5",
		  "@babel/plugin-transform-modules-commonjs": "^7.7.5",
		  "@babel/plugin-transform-modules-systemjs": "^7.7.4",
		  "@babel/plugin-transform-modules-umd": "^7.7.4",
		  "@babel/plugin-transform-named-capturing-groups-regex": "^7.7.4",
		  "@babel/plugin-transform-new-target": "^7.7.4",
		  "@babel/plugin-transform-object-super": "^7.7.4",
		  "@babel/plugin-transform-parameters": "^7.7.7",
		  "@babel/plugin-transform-property-literals": "^7.7.4",
		  "@babel/plugin-transform-regenerator": "^7.7.5",
		  "@babel/plugin-transform-reserved-words": "^7.7.4",
		  "@babel/plugin-transform-shorthand-properties": "^7.7.4",
		  "@babel/plugin-transform-spread": "^7.7.4",
		  "@babel/plugin-transform-sticky-regex": "^7.7.4",
		  "@babel/plugin-transform-template-literals": "^7.7.4",
		  "@babel/plugin-transform-typeof-symbol": "^7.7.4",
		  "@babel/plugin-transform-unicode-regex": "^7.7.4",
		  "@babel/types": "^7.7.4",
		  "browserslist": "^4.6.0",
		  "core-js-compat": "^3.6.0",
		  "invariant": "^2.2.2",
		  "js-levenshtein": "^1.1.3",
		  "semver": "^5.5.0"
		}
	  },
	  "@babel/template": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2ftemplate/-/template-7.7.4.tgz",
		"integrity": "sha512-qUzihgVPguAzXCK7WXw8pqs6cEwi54s3E+HrejlkuWO6ivMKx9hZl3Y2fSXp9i5HgyWmj7RKP+ulaYnKM4yYxw==",
		"dev": true,
		"requires": {
		  "@babel/code-frame": "^7.0.0",
		  "@babel/parser": "^7.7.4",
		  "@babel/types": "^7.7.4"
		}
	  },
	  "@babel/traverse": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2ftraverse/-/traverse-7.7.4.tgz",
		"integrity": "sha512-P1L58hQyupn8+ezVA2z5KBm4/Zr4lCC8dwKCMYzsa5jFMDMQAzaBNy9W5VjB+KAmBjb40U7a/H6ao+Xo+9saIw==",
		"dev": true,
		"requires": {
		  "@babel/code-frame": "^7.5.5",
		  "@babel/generator": "^7.7.4",
		  "@babel/helper-function-name": "^7.7.4",
		  "@babel/helper-split-export-declaration": "^7.7.4",
		  "@babel/parser": "^7.7.4",
		  "@babel/types": "^7.7.4",
		  "debug": "^4.1.0",
		  "globals": "^11.1.0",
		  "lodash": "^4.17.13"
		}
	  },
	  "@babel/types": {
		"version": "7.7.4",
		"resolved": "https://npm.intra.longguikeji.com/@babel%2ftypes/-/types-7.7.4.tgz",
		"integrity": "sha512-cz5Ji23KCi4T+YIE/BolWosrJuSmoZeN1EFnRtBwF+KKLi8GG/Z2c2hOJJeCXPk4mwk4QFvTmwIodJowXgttRA==",
		"dev": true,
		"requires": {
		  "esutils": "^2.0.2",
		  "lodash": "^4.17.13",
		  "to-fast-properties": "^2.0.0"
		}
	  },
	  "@types/color-name": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/@types%2fcolor-name/-/color-name-1.1.1.tgz",
		"integrity": "sha512-rr+OQyAjxze7GgWrSaJwydHStIhHq2lvY3BOC2Mj7KnzI7XK0Uw1TOOdI9lDoajEbSWLiYgoo4f1R51erQfhPQ=="
	  },
	  "@webassemblyjs/ast": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fast/-/ast-1.8.5.tgz",
		"integrity": "sha512-aJMfngIZ65+t71C3y2nBBg5FFG0Okt9m0XEgWZ7Ywgn1oMAT8cNwx00Uv1cQyHtidq0Xn94R4TAywO+LCQ+ZAQ==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/helper-module-context": "1.8.5",
		  "@webassemblyjs/helper-wasm-bytecode": "1.8.5",
		  "@webassemblyjs/wast-parser": "1.8.5"
		}
	  },
	  "@webassemblyjs/floating-point-hex-parser": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2ffloating-point-hex-parser/-/floating-point-hex-parser-1.8.5.tgz",
		"integrity": "sha512-9p+79WHru1oqBh9ewP9zW95E3XAo+90oth7S5Re3eQnECGq59ly1Ri5tsIipKGpiStHsUYmY3zMLqtk3gTcOtQ==",
		"dev": true
	  },
	  "@webassemblyjs/helper-api-error": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fhelper-api-error/-/helper-api-error-1.8.5.tgz",
		"integrity": "sha512-Za/tnzsvnqdaSPOUXHyKJ2XI7PDX64kWtURyGiJJZKVEdFOsdKUCPTNEVFZq3zJ2R0G5wc2PZ5gvdTRFgm81zA==",
		"dev": true
	  },
	  "@webassemblyjs/helper-buffer": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fhelper-buffer/-/helper-buffer-1.8.5.tgz",
		"integrity": "sha512-Ri2R8nOS0U6G49Q86goFIPNgjyl6+oE1abW1pS84BuhP1Qcr5JqMwRFT3Ah3ADDDYGEgGs1iyb1DGX+kAi/c/Q==",
		"dev": true
	  },
	  "@webassemblyjs/helper-code-frame": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fhelper-code-frame/-/helper-code-frame-1.8.5.tgz",
		"integrity": "sha512-VQAadSubZIhNpH46IR3yWO4kZZjMxN1opDrzePLdVKAZ+DFjkGD/rf4v1jap744uPVU6yjL/smZbRIIJTOUnKQ==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/wast-printer": "1.8.5"
		}
	  },
	  "@webassemblyjs/helper-fsm": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fhelper-fsm/-/helper-fsm-1.8.5.tgz",
		"integrity": "sha512-kRuX/saORcg8se/ft6Q2UbRpZwP4y7YrWsLXPbbmtepKr22i8Z4O3V5QE9DbZK908dh5Xya4Un57SDIKwB9eow==",
		"dev": true
	  },
	  "@webassemblyjs/helper-module-context": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fhelper-module-context/-/helper-module-context-1.8.5.tgz",
		"integrity": "sha512-/O1B236mN7UNEU4t9X7Pj38i4VoU8CcMHyy3l2cV/kIF4U5KoHXDVqcDuOs1ltkac90IM4vZdHc52t1x8Yfs3g==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/ast": "1.8.5",
		  "mamacro": "^0.0.3"
		}
	  },
	  "@webassemblyjs/helper-wasm-bytecode": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fhelper-wasm-bytecode/-/helper-wasm-bytecode-1.8.5.tgz",
		"integrity": "sha512-Cu4YMYG3Ddl72CbmpjU/wbP6SACcOPVbHN1dI4VJNJVgFwaKf1ppeFJrwydOG3NDHxVGuCfPlLZNyEdIYlQ6QQ==",
		"dev": true
	  },
	  "@webassemblyjs/helper-wasm-section": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fhelper-wasm-section/-/helper-wasm-section-1.8.5.tgz",
		"integrity": "sha512-VV083zwR+VTrIWWtgIUpqfvVdK4ff38loRmrdDBgBT8ADXYsEZ5mPQ4Nde90N3UYatHdYoDIFb7oHzMncI02tA==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/ast": "1.8.5",
		  "@webassemblyjs/helper-buffer": "1.8.5",
		  "@webassemblyjs/helper-wasm-bytecode": "1.8.5",
		  "@webassemblyjs/wasm-gen": "1.8.5"
		}
	  },
	  "@webassemblyjs/ieee754": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fieee754/-/ieee754-1.8.5.tgz",
		"integrity": "sha512-aaCvQYrvKbY/n6wKHb/ylAJr27GglahUO89CcGXMItrOBqRarUMxWLJgxm9PJNuKULwN5n1csT9bYoMeZOGF3g==",
		"dev": true,
		"requires": {
		  "@xtuc/ieee754": "^1.2.0"
		}
	  },
	  "@webassemblyjs/leb128": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fleb128/-/leb128-1.8.5.tgz",
		"integrity": "sha512-plYUuUwleLIziknvlP8VpTgO4kqNaH57Y3JnNa6DLpu/sGcP6hbVdfdX5aHAV716pQBKrfuU26BJK29qY37J7A==",
		"dev": true,
		"requires": {
		  "@xtuc/long": "4.2.2"
		}
	  },
	  "@webassemblyjs/utf8": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2futf8/-/utf8-1.8.5.tgz",
		"integrity": "sha512-U7zgftmQriw37tfD934UNInokz6yTmn29inT2cAetAsaU9YeVCveWEwhKL1Mg4yS7q//NGdzy79nlXh3bT8Kjw==",
		"dev": true
	  },
	  "@webassemblyjs/wasm-edit": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fwasm-edit/-/wasm-edit-1.8.5.tgz",
		"integrity": "sha512-A41EMy8MWw5yvqj7MQzkDjU29K7UJq1VrX2vWLzfpRHt3ISftOXqrtojn7nlPsZ9Ijhp5NwuODuycSvfAO/26Q==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/ast": "1.8.5",
		  "@webassemblyjs/helper-buffer": "1.8.5",
		  "@webassemblyjs/helper-wasm-bytecode": "1.8.5",
		  "@webassemblyjs/helper-wasm-section": "1.8.5",
		  "@webassemblyjs/wasm-gen": "1.8.5",
		  "@webassemblyjs/wasm-opt": "1.8.5",
		  "@webassemblyjs/wasm-parser": "1.8.5",
		  "@webassemblyjs/wast-printer": "1.8.5"
		}
	  },
	  "@webassemblyjs/wasm-gen": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fwasm-gen/-/wasm-gen-1.8.5.tgz",
		"integrity": "sha512-BCZBT0LURC0CXDzj5FXSc2FPTsxwp3nWcqXQdOZE4U7h7i8FqtFK5Egia6f9raQLpEKT1VL7zr4r3+QX6zArWg==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/ast": "1.8.5",
		  "@webassemblyjs/helper-wasm-bytecode": "1.8.5",
		  "@webassemblyjs/ieee754": "1.8.5",
		  "@webassemblyjs/leb128": "1.8.5",
		  "@webassemblyjs/utf8": "1.8.5"
		}
	  },
	  "@webassemblyjs/wasm-opt": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fwasm-opt/-/wasm-opt-1.8.5.tgz",
		"integrity": "sha512-HKo2mO/Uh9A6ojzu7cjslGaHaUU14LdLbGEKqTR7PBKwT6LdPtLLh9fPY33rmr5wcOMrsWDbbdCHq4hQUdd37Q==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/ast": "1.8.5",
		  "@webassemblyjs/helper-buffer": "1.8.5",
		  "@webassemblyjs/wasm-gen": "1.8.5",
		  "@webassemblyjs/wasm-parser": "1.8.5"
		}
	  },
	  "@webassemblyjs/wasm-parser": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fwasm-parser/-/wasm-parser-1.8.5.tgz",
		"integrity": "sha512-pi0SYE9T6tfcMkthwcgCpL0cM9nRYr6/6fjgDtL6q/ZqKHdMWvxitRi5JcZ7RI4SNJJYnYNaWy5UUrHQy998lw==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/ast": "1.8.5",
		  "@webassemblyjs/helper-api-error": "1.8.5",
		  "@webassemblyjs/helper-wasm-bytecode": "1.8.5",
		  "@webassemblyjs/ieee754": "1.8.5",
		  "@webassemblyjs/leb128": "1.8.5",
		  "@webassemblyjs/utf8": "1.8.5"
		}
	  },
	  "@webassemblyjs/wast-parser": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fwast-parser/-/wast-parser-1.8.5.tgz",
		"integrity": "sha512-daXC1FyKWHF1i11obK086QRlsMsY4+tIOKgBqI1lxAnkp9xe9YMcgOxm9kLe+ttjs5aWV2KKE1TWJCN57/Btsg==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/ast": "1.8.5",
		  "@webassemblyjs/floating-point-hex-parser": "1.8.5",
		  "@webassemblyjs/helper-api-error": "1.8.5",
		  "@webassemblyjs/helper-code-frame": "1.8.5",
		  "@webassemblyjs/helper-fsm": "1.8.5",
		  "@xtuc/long": "4.2.2"
		}
	  },
	  "@webassemblyjs/wast-printer": {
		"version": "1.8.5",
		"resolved": "https://npm.intra.longguikeji.com/@webassemblyjs%2fwast-printer/-/wast-printer-1.8.5.tgz",
		"integrity": "sha512-w0U0pD4EhlnvRyeJzBqaVSJAo9w/ce7/WPogeXLzGkO6hzhr4GnQIZ4W4uUt5b9ooAaXPtnXlj0gzsXEOUNYMg==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/ast": "1.8.5",
		  "@webassemblyjs/wast-parser": "1.8.5",
		  "@xtuc/long": "4.2.2"
		}
	  },
	  "@xtuc/ieee754": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/@xtuc%2fieee754/-/ieee754-1.2.0.tgz",
		"integrity": "sha512-DX8nKgqcGwsc0eJSqYt5lwP4DH5FlHnmuWWBRy7X0NcaGR0ZtuyeESgMwTYVEtxmsNGY+qit4QYT/MIYTOTPeA==",
		"dev": true
	  },
	  "@xtuc/long": {
		"version": "4.2.2",
		"resolved": "https://npm.intra.longguikeji.com/@xtuc%2flong/-/long-4.2.2.tgz",
		"integrity": "sha512-NuHqBY1PB/D8xU6s/thBgOAiAP7HOYDQ32+BFZILJ8ivkUkAHQnWfn6WhL79Owj1qmUnoN/YPhktdIoucipkAQ==",
		"dev": true
	  },
	  "accepts": {
		"version": "1.3.7",
		"resolved": "https://npm.intra.longguikeji.com/accepts/-/accepts-1.3.7.tgz",
		"integrity": "sha512-Il80Qs2WjYlJIBNzNkK6KYqlVMTbZLXgHx2oT0pU/fjRHyEp+PEfEPY0R3WCwAGVOtauxh1hOxNgIf5bv7dQpA==",
		"requires": {
		  "mime-types": "~2.1.24",
		  "negotiator": "0.6.2"
		}
	  },
	  "acorn": {
		"version": "6.4.0",
		"resolved": "https://npm.intra.longguikeji.com/acorn/-/acorn-6.4.0.tgz",
		"integrity": "sha512-gac8OEcQ2Li1dxIEWGZzsp2BitJxwkwcOm0zHAJLcPJaVvm58FRnk6RkuLRpU1EujipU2ZFODv2P9DLMfnV8mw==",
		"dev": true
	  },
	  "ajv": {
		"version": "6.10.2",
		"resolved": "https://npm.intra.longguikeji.com/ajv/-/ajv-6.10.2.tgz",
		"integrity": "sha512-TXtUUEYHuaTEbLZWIKUr5pmBuhDLy+8KYtPYdcV8qC+pOZL+NKqYwvWSRrVXHn+ZmRRAu8vJTAznH7Oag6RVRw==",
		"dev": true,
		"requires": {
		  "fast-deep-equal": "^2.0.1",
		  "fast-json-stable-stringify": "^2.0.0",
		  "json-schema-traverse": "^0.4.1",
		  "uri-js": "^4.2.2"
		}
	  },
	  "ajv-errors": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/ajv-errors/-/ajv-errors-1.0.1.tgz",
		"integrity": "sha512-DCRfO/4nQ+89p/RK43i8Ezd41EqdGIU4ld7nGF8OQ14oc/we5rEntLCUa7+jrn3nn83BosfwZA0wb4pon2o8iQ==",
		"dev": true
	  },
	  "ajv-keywords": {
		"version": "3.4.1",
		"resolved": "https://npm.intra.longguikeji.com/ajv-keywords/-/ajv-keywords-3.4.1.tgz",
		"integrity": "sha512-RO1ibKvd27e6FEShVFfPALuHI3WjSVNeK5FIsmme/LYRNxjKuNj+Dt7bucLa6NdSv3JcVTyMlm9kGR84z1XpaQ==",
		"dev": true
	  },
	  "ansi-regex": {
		"version": "5.0.0",
		"resolved": "https://npm.intra.longguikeji.com/ansi-regex/-/ansi-regex-5.0.0.tgz",
		"integrity": "sha512-bY6fj56OUQ0hU1KjFNDQuJFezqKdrAyFdIevADiqrWHwSlbmBNMHp5ak2f40Pm8JTFyM2mqxkG6ngkHO11f/lg=="
	  },
	  "ansi-styles": {
		"version": "4.2.1",
		"resolved": "https://npm.intra.longguikeji.com/ansi-styles/-/ansi-styles-4.2.1.tgz",
		"integrity": "sha512-9VGjrMsG1vePxcSweQsN20KY/c4zN0h9fLjqAbwbPfahM3t+NL+M9HC8xeXG2I8pX5NoamTGNuomEUFI7fcUjA==",
		"requires": {
		  "@types/color-name": "^1.1.1",
		  "color-convert": "^2.0.1"
		}
	  },
	  "anymatch": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/anymatch/-/anymatch-2.0.0.tgz",
		"integrity": "sha512-5teOsQWABXHHBFP9y3skS5P3d/WfWXpv3FUpy+LorMrNYaT9pI4oLMQX7jzQ2KklNpGpWHzdCXTDT2Y3XGlZBw==",
		"dev": true,
		"requires": {
		  "micromatch": "^3.1.4",
		  "normalize-path": "^2.1.1"
		},
		"dependencies": {
		  "normalize-path": {
			"version": "2.1.1",
			"resolved": "https://npm.intra.longguikeji.com/normalize-path/-/normalize-path-2.1.1.tgz",
			"integrity": "sha1-GrKLVW4Zg2Oowab35vogE3/mrtk=",
			"dev": true,
			"requires": {
			  "remove-trailing-separator": "^1.0.1"
			}
		  }
		}
	  },
	  "aproba": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/aproba/-/aproba-1.2.0.tgz",
		"integrity": "sha512-Y9J6ZjXtoYh8RnXVCMOU/ttDmk1aBjunq9vO0ta5x85WDQiQfUF9sIPBITdbiiIVcBo03Hi3jMxigBtsddlXRw==",
		"dev": true
	  },
	  "arkfbp": {
		"version": "0.0.11",
		"resolved": "https://npm.intra.longguikeji.com/arkfbp/-/arkfbp-0.0.11.tgz",
		"integrity": "sha512-pUC6McAWy2z0jusWQudRvm/XFU7ux8HxZy1mFId45aJB8QLoXNpMNMnrlM2mrE0fnqnAgwg47lvc3Nz+7CZhsA==",
		"requires": {
		  "axios": "^0.19.0",
		  "cookie": "^0.4.0",
		  "debug": "^4.1.1",
		  "express": "^4.17.1",
		  "loadash": "^1.0.0"
		}
	  },
	  "arr-diff": {
		"version": "4.0.0",
		"resolved": "https://npm.intra.longguikeji.com/arr-diff/-/arr-diff-4.0.0.tgz",
		"integrity": "sha1-1kYQdP6/7HHn4VI1dhoyml3HxSA=",
		"dev": true
	  },
	  "arr-flatten": {
		"version": "1.1.0",
		"resolved": "https://npm.intra.longguikeji.com/arr-flatten/-/arr-flatten-1.1.0.tgz",
		"integrity": "sha512-L3hKV5R/p5o81R7O02IGnwpDmkp6E982XhtbuwSe3O4qOtMMMtodicASA1Cny2U+aCXcNpml+m4dPsvsJ3jatg==",
		"dev": true
	  },
	  "arr-union": {
		"version": "3.1.0",
		"resolved": "https://npm.intra.longguikeji.com/arr-union/-/arr-union-3.1.0.tgz",
		"integrity": "sha1-45sJrqne+Gao8gbiiK9jkZuuOcQ=",
		"dev": true
	  },
	  "array-flatten": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/array-flatten/-/array-flatten-1.1.1.tgz",
		"integrity": "sha1-ml9pkFGx5wczKPKgCJaLZOopVdI="
	  },
	  "array-unique": {
		"version": "0.3.2",
		"resolved": "https://npm.intra.longguikeji.com/array-unique/-/array-unique-0.3.2.tgz",
		"integrity": "sha1-qJS3XUvE9s1nnvMkSp/Y9Gri1Cg=",
		"dev": true
	  },
	  "asn1.js": {
		"version": "4.10.1",
		"resolved": "https://npm.intra.longguikeji.com/asn1.js/-/asn1.js-4.10.1.tgz",
		"integrity": "sha512-p32cOF5q0Zqs9uBiONKYLm6BClCoBCM5O9JfeUSlnQLBTxYdTK+pW+nXflm8UkKd2UYlEbYz5qEi0JuZR9ckSw==",
		"dev": true,
		"requires": {
		  "bn.js": "^4.0.0",
		  "inherits": "^2.0.1",
		  "minimalistic-assert": "^1.0.0"
		}
	  },
	  "assert": {
		"version": "1.5.0",
		"resolved": "https://npm.intra.longguikeji.com/assert/-/assert-1.5.0.tgz",
		"integrity": "sha512-EDsgawzwoun2CZkCgtxJbv392v4nbk9XDD06zI+kQYoBM/3RBWLlEyJARDOmhAAosBjWACEkKL6S+lIZtcAubA==",
		"dev": true,
		"requires": {
		  "object-assign": "^4.1.1",
		  "util": "0.10.3"
		},
		"dependencies": {
		  "inherits": {
			"version": "2.0.1",
			"resolved": "https://npm.intra.longguikeji.com/inherits/-/inherits-2.0.1.tgz",
			"integrity": "sha1-sX0I0ya0Qj5Wjv9xn5GwscvfafE=",
			"dev": true
		  },
		  "util": {
			"version": "0.10.3",
			"resolved": "https://npm.intra.longguikeji.com/util/-/util-0.10.3.tgz",
			"integrity": "sha1-evsa/lCAUkZInj23/g7TeTNqwPk=",
			"dev": true,
			"requires": {
			  "inherits": "2.0.1"
			}
		  }
		}
	  },
	  "assign-symbols": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/assign-symbols/-/assign-symbols-1.0.0.tgz",
		"integrity": "sha1-WWZ/QfrdTyDMvCu5a41Pf3jsA2c=",
		"dev": true
	  },
	  "async-each": {
		"version": "1.0.3",
		"resolved": "https://npm.intra.longguikeji.com/async-each/-/async-each-1.0.3.tgz",
		"integrity": "sha512-z/WhQ5FPySLdvREByI2vZiTWwCnF0moMJ1hK9YQwDTHKh6I7/uSckMetoRGb5UBZPC1z0jlw+n/XCgjeH7y1AQ==",
		"dev": true
	  },
	  "atob": {
		"version": "2.1.2",
		"resolved": "https://npm.intra.longguikeji.com/atob/-/atob-2.1.2.tgz",
		"integrity": "sha512-Wm6ukoaOGJi/73p/cl2GvLjTI5JM1k/O14isD73YML8StrH/7/lRFgmg8nICZgD3bZZvjwCGxtMOD3wWNAu8cg==",
		"dev": true
	  },
	  "axios": {
		"version": "0.19.0",
		"resolved": "https://npm.intra.longguikeji.com/axios/-/axios-0.19.0.tgz",
		"integrity": "sha512-1uvKqKQta3KBxIz14F2v06AEHZ/dIoeKfbTRkK1E5oqjDnuEerLmYTgJB5AiQZHJcljpg1TuRzdjDR06qNk0DQ==",
		"requires": {
		  "follow-redirects": "1.5.10",
		  "is-buffer": "^2.0.2"
		}
	  },
	  "babel-loader": {
		"version": "8.0.6",
		"resolved": "https://npm.intra.longguikeji.com/babel-loader/-/babel-loader-8.0.6.tgz",
		"integrity": "sha512-4BmWKtBOBm13uoUwd08UwjZlaw3O9GWf456R9j+5YykFZ6LUIjIKLc0zEZf+hauxPOJs96C8k6FvYD09vWzhYw==",
		"dev": true,
		"requires": {
		  "find-cache-dir": "^2.0.0",
		  "loader-utils": "^1.0.2",
		  "mkdirp": "^0.5.1",
		  "pify": "^4.0.1"
		}
	  },
	  "babel-plugin-dynamic-import-node": {
		"version": "2.3.0",
		"resolved": "https://npm.intra.longguikeji.com/babel-plugin-dynamic-import-node/-/babel-plugin-dynamic-import-node-2.3.0.tgz",
		"integrity": "sha512-o6qFkpeQEBxcqt0XYlWzAVxNCSCZdUgcR8IRlhD/8DylxjjO4foPcvTW0GGKa/cVt3rvxZ7o5ippJ+/0nvLhlQ==",
		"dev": true,
		"requires": {
		  "object.assign": "^4.1.0"
		}
	  },
	  "balanced-match": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/balanced-match/-/balanced-match-1.0.0.tgz",
		"integrity": "sha1-ibTRmasr7kneFk6gK4nORi1xt2c=",
		"dev": true
	  },
	  "base": {
		"version": "0.11.2",
		"resolved": "https://npm.intra.longguikeji.com/base/-/base-0.11.2.tgz",
		"integrity": "sha512-5T6P4xPgpp0YDFvSWwEZ4NoE3aM4QBQXDzmVbraCkFj8zHM+mba8SyqB5DbZWyR7mYHo6Y7BdQo3MoA4m0TeQg==",
		"dev": true,
		"requires": {
		  "cache-base": "^1.0.1",
		  "class-utils": "^0.3.5",
		  "component-emitter": "^1.2.1",
		  "define-property": "^1.0.0",
		  "isobject": "^3.0.1",
		  "mixin-deep": "^1.2.0",
		  "pascalcase": "^0.1.1"
		},
		"dependencies": {
		  "define-property": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/define-property/-/define-property-1.0.0.tgz",
			"integrity": "sha1-dp66rz9KY6rTr56NMEybvnm/sOY=",
			"dev": true,
			"requires": {
			  "is-descriptor": "^1.0.0"
			}
		  },
		  "is-accessor-descriptor": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/is-accessor-descriptor/-/is-accessor-descriptor-1.0.0.tgz",
			"integrity": "sha512-m5hnHTkcVsPfqx3AKlyttIPb7J+XykHvJP2B9bZDjlhLIoEq4XoK64Vg7boZlVWYK6LUY94dYPEE7Lh0ZkZKcQ==",
			"dev": true,
			"requires": {
			  "kind-of": "^6.0.0"
			}
		  },
		  "is-data-descriptor": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/is-data-descriptor/-/is-data-descriptor-1.0.0.tgz",
			"integrity": "sha512-jbRXy1FmtAoCjQkVmIVYwuuqDFUbaOeDjmed1tOGPrsMhtJA4rD9tkgA0F1qJ3gRFRXcHYVkdeaP50Q5rE/jLQ==",
			"dev": true,
			"requires": {
			  "kind-of": "^6.0.0"
			}
		  },
		  "is-descriptor": {
			"version": "1.0.2",
			"resolved": "https://npm.intra.longguikeji.com/is-descriptor/-/is-descriptor-1.0.2.tgz",
			"integrity": "sha512-2eis5WqQGV7peooDyLmNEPUrps9+SXX5c9pL3xEB+4e9HnGuDa7mB7kHxHw4CbqS9k1T2hOH3miL8n8WtiYVtg==",
			"dev": true,
			"requires": {
			  "is-accessor-descriptor": "^1.0.0",
			  "is-data-descriptor": "^1.0.0",
			  "kind-of": "^6.0.2"
			}
		  }
		}
	  },
	  "base64-js": {
		"version": "1.3.1",
		"resolved": "https://npm.intra.longguikeji.com/base64-js/-/base64-js-1.3.1.tgz",
		"integrity": "sha512-mLQ4i2QO1ytvGWFWmcngKO//JXAQueZvwEKtjgQFM4jIK0kU+ytMfplL8j+n5mspOfjHwoAg+9yhb7BwAHm36g==",
		"dev": true
	  },
	  "big.js": {
		"version": "5.2.2",
		"resolved": "https://npm.intra.longguikeji.com/big.js/-/big.js-5.2.2.tgz",
		"integrity": "sha512-vyL2OymJxmarO8gxMr0mhChsO9QGwhynfuu4+MHTAW6czfq9humCB7rKpUjDd9YUiDPU4mzpyupFSvOClAwbmQ==",
		"dev": true
	  },
	  "binary-extensions": {
		"version": "1.13.1",
		"resolved": "https://npm.intra.longguikeji.com/binary-extensions/-/binary-extensions-1.13.1.tgz",
		"integrity": "sha512-Un7MIEDdUC5gNpcGDV97op1Ywk748MpHcFTHoYs6qnj1Z3j7I53VG3nwZhKzoBZmbdRNnb6WRdFlwl7tSDuZGw==",
		"dev": true
	  },
	  "bindings": {
		"version": "1.5.0",
		"resolved": "https://npm.intra.longguikeji.com/bindings/-/bindings-1.5.0.tgz",
		"integrity": "sha512-p2q/t/mhvuOj/UeLlV6566GD/guowlr0hHxClI0W9m7MWYkL1F0hLo+0Aexs9HSPCtR1SXQ0TD3MMKrXZajbiQ==",
		"dev": true,
		"optional": true,
		"requires": {
		  "file-uri-to-path": "1.0.0"
		}
	  },
	  "bluebird": {
		"version": "3.7.2",
		"resolved": "https://npm.intra.longguikeji.com/bluebird/-/bluebird-3.7.2.tgz",
		"integrity": "sha512-XpNj6GDQzdfW+r2Wnn7xiSAd7TM3jzkxGXBGTtWKuSXv1xUV+azxAm8jdWZN06QTQk+2N2XB9jRDkvbmQmcRtg==",
		"dev": true
	  },
	  "bn.js": {
		"version": "4.11.8",
		"resolved": "https://npm.intra.longguikeji.com/bn.js/-/bn.js-4.11.8.tgz",
		"integrity": "sha512-ItfYfPLkWHUjckQCk8xC+LwxgK8NYcXywGigJgSwOP8Y2iyWT4f2vsZnoOXTTbo+o5yXmIUJ4gn5538SO5S3gA==",
		"dev": true
	  },
	  "body-parser": {
		"version": "1.19.0",
		"resolved": "https://npm.intra.longguikeji.com/body-parser/-/body-parser-1.19.0.tgz",
		"integrity": "sha512-dhEPs72UPbDnAQJ9ZKMNTP6ptJaionhP5cBb541nXPlW60Jepo9RV/a4fX4XWW9CuFNK22krhrj1+rgzifNCsw==",
		"requires": {
		  "bytes": "3.1.0",
		  "content-type": "~1.0.4",
		  "debug": "2.6.9",
		  "depd": "~1.1.2",
		  "http-errors": "1.7.2",
		  "iconv-lite": "0.4.24",
		  "on-finished": "~2.3.0",
		  "qs": "6.7.0",
		  "raw-body": "2.4.0",
		  "type-is": "~1.6.17"
		},
		"dependencies": {
		  "debug": {
			"version": "2.6.9",
			"resolved": "https://npm.intra.longguikeji.com/debug/-/debug-2.6.9.tgz",
			"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
			"requires": {
			  "ms": "2.0.0"
			}
		  }
		}
	  },
	  "brace-expansion": {
		"version": "1.1.11",
		"resolved": "https://npm.intra.longguikeji.com/brace-expansion/-/brace-expansion-1.1.11.tgz",
		"integrity": "sha512-iCuPHDFgrHX7H2vEI/5xpz07zSHB00TpugqhmYtVmMO6518mCuRMoOYFldEBl0g187ufozdaHgWKcYFb61qGiA==",
		"dev": true,
		"requires": {
		  "balanced-match": "^1.0.0",
		  "concat-map": "0.0.1"
		}
	  },
	  "braces": {
		"version": "2.3.2",
		"resolved": "https://npm.intra.longguikeji.com/braces/-/braces-2.3.2.tgz",
		"integrity": "sha512-aNdbnj9P8PjdXU4ybaWLK2IF3jc/EoDYbC7AazW6to3TRsfXxscC9UXOB5iDiEQrkyIbWp2SLQda4+QAa7nc3w==",
		"dev": true,
		"requires": {
		  "arr-flatten": "^1.1.0",
		  "array-unique": "^0.3.2",
		  "extend-shallow": "^2.0.1",
		  "fill-range": "^4.0.0",
		  "isobject": "^3.0.1",
		  "repeat-element": "^1.1.2",
		  "snapdragon": "^0.8.1",
		  "snapdragon-node": "^2.0.1",
		  "split-string": "^3.0.2",
		  "to-regex": "^3.0.1"
		},
		"dependencies": {
		  "extend-shallow": {
			"version": "2.0.1",
			"resolved": "https://npm.intra.longguikeji.com/extend-shallow/-/extend-shallow-2.0.1.tgz",
			"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
			"dev": true,
			"requires": {
			  "is-extendable": "^0.1.0"
			}
		  }
		}
	  },
	  "brorand": {
		"version": "1.1.0",
		"resolved": "https://npm.intra.longguikeji.com/brorand/-/brorand-1.1.0.tgz",
		"integrity": "sha1-EsJe/kCkXjwyPrhnWgoM5XsiNx8=",
		"dev": true
	  },
	  "browserify-aes": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/browserify-aes/-/browserify-aes-1.2.0.tgz",
		"integrity": "sha512-+7CHXqGuspUn/Sl5aO7Ea0xWGAtETPXNSAjHo48JfLdPWcMng33Xe4znFvQweqc/uzk5zSOI3H52CYnjCfb5hA==",
		"dev": true,
		"requires": {
		  "buffer-xor": "^1.0.3",
		  "cipher-base": "^1.0.0",
		  "create-hash": "^1.1.0",
		  "evp_bytestokey": "^1.0.3",
		  "inherits": "^2.0.1",
		  "safe-buffer": "^5.0.1"
		}
	  },
	  "browserify-cipher": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/browserify-cipher/-/browserify-cipher-1.0.1.tgz",
		"integrity": "sha512-sPhkz0ARKbf4rRQt2hTpAHqn47X3llLkUGn+xEJzLjwY8LRs2p0v7ljvI5EyoRO/mexrNunNECisZs+gw2zz1w==",
		"dev": true,
		"requires": {
		  "browserify-aes": "^1.0.4",
		  "browserify-des": "^1.0.0",
		  "evp_bytestokey": "^1.0.0"
		}
	  },
	  "browserify-des": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/browserify-des/-/browserify-des-1.0.2.tgz",
		"integrity": "sha512-BioO1xf3hFwz4kc6iBhI3ieDFompMhrMlnDFC4/0/vd5MokpuAc3R+LYbwTA9A5Yc9pq9UYPqffKpW2ObuwX5A==",
		"dev": true,
		"requires": {
		  "cipher-base": "^1.0.1",
		  "des.js": "^1.0.0",
		  "inherits": "^2.0.1",
		  "safe-buffer": "^5.1.2"
		}
	  },
	  "browserify-rsa": {
		"version": "4.0.1",
		"resolved": "https://npm.intra.longguikeji.com/browserify-rsa/-/browserify-rsa-4.0.1.tgz",
		"integrity": "sha1-IeCr+vbyApzy+vsTNWenAdQTVSQ=",
		"dev": true,
		"requires": {
		  "bn.js": "^4.1.0",
		  "randombytes": "^2.0.1"
		}
	  },
	  "browserify-sign": {
		"version": "4.0.4",
		"resolved": "https://npm.intra.longguikeji.com/browserify-sign/-/browserify-sign-4.0.4.tgz",
		"integrity": "sha1-qk62jl17ZYuqa/alfmMMvXqT0pg=",
		"dev": true,
		"requires": {
		  "bn.js": "^4.1.1",
		  "browserify-rsa": "^4.0.0",
		  "create-hash": "^1.1.0",
		  "create-hmac": "^1.1.2",
		  "elliptic": "^6.0.0",
		  "inherits": "^2.0.1",
		  "parse-asn1": "^5.0.0"
		}
	  },
	  "browserify-zlib": {
		"version": "0.2.0",
		"resolved": "https://npm.intra.longguikeji.com/browserify-zlib/-/browserify-zlib-0.2.0.tgz",
		"integrity": "sha512-Z942RysHXmJrhqk88FmKBVq/v5tqmSkDz7p54G/MGyjMnCFFnC79XWNbg+Vta8W6Wb2qtSZTSxIGkJrRpCFEiA==",
		"dev": true,
		"requires": {
		  "pako": "~1.0.5"
		}
	  },
	  "browserslist": {
		"version": "4.8.3",
		"resolved": "https://npm.intra.longguikeji.com/browserslist/-/browserslist-4.8.3.tgz",
		"integrity": "sha512-iU43cMMknxG1ClEZ2MDKeonKE1CCrFVkQK2AqO2YWFmvIrx4JWrvQ4w4hQez6EpVI8rHTtqh/ruHHDHSOKxvUg==",
		"dev": true,
		"requires": {
		  "caniuse-lite": "^1.0.30001017",
		  "electron-to-chromium": "^1.3.322",
		  "node-releases": "^1.1.44"
		}
	  },
	  "buffer": {
		"version": "4.9.2",
		"resolved": "https://npm.intra.longguikeji.com/buffer/-/buffer-4.9.2.tgz",
		"integrity": "sha512-xq+q3SRMOxGivLhBNaUdC64hDTQwejJ+H0T/NB1XMtTVEwNTrfFF3gAxiyW0Bu/xWEGhjVKgUcMhCrUy2+uCWg==",
		"dev": true,
		"requires": {
		  "base64-js": "^1.0.2",
		  "ieee754": "^1.1.4",
		  "isarray": "^1.0.0"
		}
	  },
	  "buffer-from": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/buffer-from/-/buffer-from-1.1.1.tgz",
		"integrity": "sha512-MQcXEUbCKtEo7bhqEs6560Hyd4XaovZlO/k9V3hjVUF/zwW7KBVdSK4gIt/bzwS9MbR5qob+F5jusZsb0YQK2A==",
		"dev": true
	  },
	  "buffer-xor": {
		"version": "1.0.3",
		"resolved": "https://npm.intra.longguikeji.com/buffer-xor/-/buffer-xor-1.0.3.tgz",
		"integrity": "sha1-JuYe0UIvtw3ULm42cp7VHYVf6Nk=",
		"dev": true
	  },
	  "builtin-status-codes": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/builtin-status-codes/-/builtin-status-codes-3.0.0.tgz",
		"integrity": "sha1-hZgoeOIbmOHGZCXgPQF0eI9Wnug=",
		"dev": true
	  },
	  "bytes": {
		"version": "3.1.0",
		"resolved": "https://npm.intra.longguikeji.com/bytes/-/bytes-3.1.0.tgz",
		"integrity": "sha512-zauLjrfCG+xvoyaqLoV8bLVXXNGC4JqlxFCutSDWA6fJrTo2ZuvLYTqZ7aHBLZSMOopbzwv8f+wZcVzfVTI2Dg=="
	  },
	  "cacache": {
		"version": "12.0.3",
		"resolved": "https://npm.intra.longguikeji.com/cacache/-/cacache-12.0.3.tgz",
		"integrity": "sha512-kqdmfXEGFepesTuROHMs3MpFLWrPkSSpRqOw80RCflZXy/khxaArvFrQ7uJxSUduzAufc6G0g1VUCOZXxWavPw==",
		"dev": true,
		"requires": {
		  "bluebird": "^3.5.5",
		  "chownr": "^1.1.1",
		  "figgy-pudding": "^3.5.1",
		  "glob": "^7.1.4",
		  "graceful-fs": "^4.1.15",
		  "infer-owner": "^1.0.3",
		  "lru-cache": "^5.1.1",
		  "mississippi": "^3.0.0",
		  "mkdirp": "^0.5.1",
		  "move-concurrently": "^1.0.1",
		  "promise-inflight": "^1.0.1",
		  "rimraf": "^2.6.3",
		  "ssri": "^6.0.1",
		  "unique-filename": "^1.1.1",
		  "y18n": "^4.0.0"
		}
	  },
	  "cache-base": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/cache-base/-/cache-base-1.0.1.tgz",
		"integrity": "sha512-AKcdTnFSWATd5/GCPRxr2ChwIJ85CeyrEyjRHlKxQ56d4XJMGym0uAiKn0xbLOGOl3+yRpOTi484dVCEc5AUzQ==",
		"dev": true,
		"requires": {
		  "collection-visit": "^1.0.0",
		  "component-emitter": "^1.2.1",
		  "get-value": "^2.0.6",
		  "has-value": "^1.0.0",
		  "isobject": "^3.0.1",
		  "set-value": "^2.0.0",
		  "to-object-path": "^0.3.0",
		  "union-value": "^1.0.0",
		  "unset-value": "^1.0.0"
		}
	  },
	  "camelcase": {
		"version": "5.3.1",
		"resolved": "https://npm.intra.longguikeji.com/camelcase/-/camelcase-5.3.1.tgz",
		"integrity": "sha512-L28STB170nwWS63UjtlEOE3dldQApaJXZkOI1uMFfzf3rRuPegHaHesyee+YxQ+W6SvRDQV6UrdOdRiR153wJg=="
	  },
	  "caniuse-lite": {
		"version": "1.0.30001019",
		"resolved": "https://npm.intra.longguikeji.com/caniuse-lite/-/caniuse-lite-1.0.30001019.tgz",
		"integrity": "sha512-6ljkLtF1KM5fQ+5ZN0wuyVvvebJxgJPTmScOMaFuQN2QuOzvRJnWSKfzQskQU5IOU4Gap3zasYPIinzwUjoj/g==",
		"dev": true
	  },
	  "chalk": {
		"version": "2.4.2",
		"resolved": "https://npm.intra.longguikeji.com/chalk/-/chalk-2.4.2.tgz",
		"integrity": "sha512-Mti+f9lpJNcwF4tWV8/OrTTtF1gZi+f8FqlyAdouralcFWFQWF2+NgCHShjkCb+IFBLq9buZwE1xckQU4peSuQ==",
		"dev": true,
		"requires": {
		  "ansi-styles": "^3.2.1",
		  "escape-string-regexp": "^1.0.5",
		  "supports-color": "^5.3.0"
		},
		"dependencies": {
		  "ansi-styles": {
			"version": "3.2.1",
			"resolved": "https://npm.intra.longguikeji.com/ansi-styles/-/ansi-styles-3.2.1.tgz",
			"integrity": "sha512-VT0ZI6kZRdTh8YyJw3SMbYm/u+NqfsAxEpWO0Pf9sq8/e94WxxOpPKx9FR1FlyCtOVDNOQ+8ntlqFxiRc+r5qA==",
			"dev": true,
			"requires": {
			  "color-convert": "^1.9.0"
			}
		  },
		  "color-convert": {
			"version": "1.9.3",
			"resolved": "https://npm.intra.longguikeji.com/color-convert/-/color-convert-1.9.3.tgz",
			"integrity": "sha512-QfAUtd+vFdAtFQcC8CCyYt1fYWxSqAiK2cSD6zDB8N3cpsEBAvRxp9zOGg6G/SHHJYAT88/az/IuDGALsNVbGg==",
			"dev": true,
			"requires": {
			  "color-name": "1.1.3"
			}
		  },
		  "color-name": {
			"version": "1.1.3",
			"resolved": "https://npm.intra.longguikeji.com/color-name/-/color-name-1.1.3.tgz",
			"integrity": "sha1-p9BVi9icQveV3UIyj3QIMcpTvCU=",
			"dev": true
		  }
		}
	  },
	  "chokidar": {
		"version": "2.1.8",
		"resolved": "https://npm.intra.longguikeji.com/chokidar/-/chokidar-2.1.8.tgz",
		"integrity": "sha512-ZmZUazfOzf0Nve7duiCKD23PFSCs4JPoYyccjUFF3aQkQadqBhfzhjkwBH2mNOG9cTBwhamM37EIsIkZw3nRgg==",
		"dev": true,
		"requires": {
		  "anymatch": "^2.0.0",
		  "async-each": "^1.0.1",
		  "braces": "^2.3.2",
		  "fsevents": "^1.2.7",
		  "glob-parent": "^3.1.0",
		  "inherits": "^2.0.3",
		  "is-binary-path": "^1.0.0",
		  "is-glob": "^4.0.0",
		  "normalize-path": "^3.0.0",
		  "path-is-absolute": "^1.0.0",
		  "readdirp": "^2.2.1",
		  "upath": "^1.1.1"
		}
	  },
	  "chownr": {
		"version": "1.1.3",
		"resolved": "https://npm.intra.longguikeji.com/chownr/-/chownr-1.1.3.tgz",
		"integrity": "sha512-i70fVHhmV3DtTl6nqvZOnIjbY0Pe4kAUjwHj8z0zAdgBtYrJyYwLKCCuRBQ5ppkyL0AkN7HKRnETdmdp1zqNXw==",
		"dev": true
	  },
	  "chrome-trace-event": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/chrome-trace-event/-/chrome-trace-event-1.0.2.tgz",
		"integrity": "sha512-9e/zx1jw7B4CO+c/RXoCsfg/x1AfUBioy4owYH0bJprEYAx5hRFLRhWBqHAG57D0ZM4H7vxbP7bPe0VwhQRYDQ==",
		"dev": true,
		"requires": {
		  "tslib": "^1.9.0"
		}
	  },
	  "cipher-base": {
		"version": "1.0.4",
		"resolved": "https://npm.intra.longguikeji.com/cipher-base/-/cipher-base-1.0.4.tgz",
		"integrity": "sha512-Kkht5ye6ZGmwv40uUDZztayT2ThLQGfnj/T71N/XzeZeo3nf8foyW7zGTsPYkEya3m5f3cAypH+qe7YOrM1U2Q==",
		"dev": true,
		"requires": {
		  "inherits": "^2.0.1",
		  "safe-buffer": "^5.0.1"
		}
	  },
	  "class-utils": {
		"version": "0.3.6",
		"resolved": "https://npm.intra.longguikeji.com/class-utils/-/class-utils-0.3.6.tgz",
		"integrity": "sha512-qOhPa/Fj7s6TY8H8esGu5QNpMMQxz79h+urzrNYN6mn+9BnxlDGf5QZ+XeCDsxSjPqsSR56XOZOJmpeurnLMeg==",
		"dev": true,
		"requires": {
		  "arr-union": "^3.1.0",
		  "define-property": "^0.2.5",
		  "isobject": "^3.0.0",
		  "static-extend": "^0.1.1"
		},
		"dependencies": {
		  "define-property": {
			"version": "0.2.5",
			"resolved": "https://npm.intra.longguikeji.com/define-property/-/define-property-0.2.5.tgz",
			"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
			"dev": true,
			"requires": {
			  "is-descriptor": "^0.1.0"
			}
		  }
		}
	  },
	  "cliui": {
		"version": "6.0.0",
		"resolved": "https://npm.intra.longguikeji.com/cliui/-/cliui-6.0.0.tgz",
		"integrity": "sha512-t6wbgtoCXvAzst7QgXxJYqPt0usEfbgQdftEPbLL/cvv6HPE5VgvqCuAIDR0NgU52ds6rFwqrgakNLrHEjCbrQ==",
		"requires": {
		  "string-width": "^4.2.0",
		  "strip-ansi": "^6.0.0",
		  "wrap-ansi": "^6.2.0"
		}
	  },
	  "collection-visit": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/collection-visit/-/collection-visit-1.0.0.tgz",
		"integrity": "sha1-S8A3PBZLwykbTTaMgpzxqApZ3KA=",
		"dev": true,
		"requires": {
		  "map-visit": "^1.0.0",
		  "object-visit": "^1.0.0"
		}
	  },
	  "color-convert": {
		"version": "2.0.1",
		"resolved": "https://npm.intra.longguikeji.com/color-convert/-/color-convert-2.0.1.tgz",
		"integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
		"requires": {
		  "color-name": "~1.1.4"
		}
	  },
	  "color-name": {
		"version": "1.1.4",
		"resolved": "https://npm.intra.longguikeji.com/color-name/-/color-name-1.1.4.tgz",
		"integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA=="
	  },
	  "commander": {
		"version": "4.1.0",
		"resolved": "https://npm.intra.longguikeji.com/commander/-/commander-4.1.0.tgz",
		"integrity": "sha512-NIQrwvv9V39FHgGFm36+U9SMQzbiHvU79k+iADraJTpmrFFfx7Ds0IvDoAdZsDrknlkRk14OYoWXb57uTh7/sw==",
		"dev": true
	  },
	  "commondir": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/commondir/-/commondir-1.0.1.tgz",
		"integrity": "sha1-3dgA2gxmEnOTzKWVDqloo6rxJTs=",
		"dev": true
	  },
	  "component-emitter": {
		"version": "1.3.0",
		"resolved": "https://npm.intra.longguikeji.com/component-emitter/-/component-emitter-1.3.0.tgz",
		"integrity": "sha512-Rd3se6QB+sO1TwqZjscQrurpEPIfO0/yYnSin6Q/rD3mOutHvUrCAhJub3r90uNb+SESBuE0QYoB90YdfatsRg==",
		"dev": true
	  },
	  "concat-map": {
		"version": "0.0.1",
		"resolved": "https://npm.intra.longguikeji.com/concat-map/-/concat-map-0.0.1.tgz",
		"integrity": "sha1-2Klr13/Wjfd5OnMDajug1UBdR3s=",
		"dev": true
	  },
	  "concat-stream": {
		"version": "1.6.2",
		"resolved": "https://npm.intra.longguikeji.com/concat-stream/-/concat-stream-1.6.2.tgz",
		"integrity": "sha512-27HBghJxjiZtIk3Ycvn/4kbJk/1uZuJFfuPEns6LaEvpvG1f0hTea8lilrouyo9mVc2GWdcEZ8OLoGmSADlrCw==",
		"dev": true,
		"requires": {
		  "buffer-from": "^1.0.0",
		  "inherits": "^2.0.3",
		  "readable-stream": "^2.2.2",
		  "typedarray": "^0.0.6"
		}
	  },
	  "console-browserify": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/console-browserify/-/console-browserify-1.2.0.tgz",
		"integrity": "sha512-ZMkYO/LkF17QvCPqM0gxw8yUzigAOZOSWSHg91FH6orS7vcEj5dVZTidN2fQ14yBSdg97RqhSNwLUXInd52OTA==",
		"dev": true
	  },
	  "constants-browserify": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/constants-browserify/-/constants-browserify-1.0.0.tgz",
		"integrity": "sha1-wguW2MYXdIqvHBYCF2DNJ/y4y3U=",
		"dev": true
	  },
	  "content-disposition": {
		"version": "0.5.3",
		"resolved": "https://npm.intra.longguikeji.com/content-disposition/-/content-disposition-0.5.3.tgz",
		"integrity": "sha512-ExO0774ikEObIAEV9kDo50o+79VCUdEB6n6lzKgGwupcVeRlhrj3qGAfwq8G6uBJjkqLrhT0qEYFcWng8z1z0g==",
		"requires": {
		  "safe-buffer": "5.1.2"
		}
	  },
	  "content-type": {
		"version": "1.0.4",
		"resolved": "https://npm.intra.longguikeji.com/content-type/-/content-type-1.0.4.tgz",
		"integrity": "sha512-hIP3EEPs8tB9AT1L+NUqtwOAps4mk2Zob89MWXMHjHWg9milF/j4osnnQLXBCBFBk/tvIG/tUc9mOUJiPBhPXA=="
	  },
	  "convert-source-map": {
		"version": "1.7.0",
		"resolved": "https://npm.intra.longguikeji.com/convert-source-map/-/convert-source-map-1.7.0.tgz",
		"integrity": "sha512-4FJkXzKXEDB1snCFZlLP4gpC3JILicCpGbzG9f9G7tGqGCzETQ2hWPrcinA9oU4wtf2biUaEH5065UnMeR33oA==",
		"dev": true,
		"requires": {
		  "safe-buffer": "~5.1.1"
		}
	  },
	  "cookie": {
		"version": "0.4.0",
		"resolved": "https://npm.intra.longguikeji.com/cookie/-/cookie-0.4.0.tgz",
		"integrity": "sha512-+Hp8fLp57wnUSt0tY0tHEXh4voZRDnoIrZPqlo3DPiI4y9lwg/jqx+1Om94/W6ZaPDOUbnjOt/99w66zk+l1Xg=="
	  },
	  "cookie-parser": {
		"version": "1.4.4",
		"resolved": "https://npm.intra.longguikeji.com/cookie-parser/-/cookie-parser-1.4.4.tgz",
		"integrity": "sha512-lo13tqF3JEtFO7FyA49CqbhaFkskRJ0u/UAiINgrIXeRCY41c88/zxtrECl8AKH3B0hj9q10+h3Kt8I7KlW4tw==",
		"requires": {
		  "cookie": "0.3.1",
		  "cookie-signature": "1.0.6"
		},
		"dependencies": {
		  "cookie": {
			"version": "0.3.1",
			"resolved": "https://npm.intra.longguikeji.com/cookie/-/cookie-0.3.1.tgz",
			"integrity": "sha1-5+Ch+e9DtMi6klxcWpboBtFoc7s="
		  }
		}
	  },
	  "cookie-signature": {
		"version": "1.0.6",
		"resolved": "https://npm.intra.longguikeji.com/cookie-signature/-/cookie-signature-1.0.6.tgz",
		"integrity": "sha1-4wOogrNCzD7oylE6eZmXNNqzriw="
	  },
	  "copy-concurrently": {
		"version": "1.0.5",
		"resolved": "https://npm.intra.longguikeji.com/copy-concurrently/-/copy-concurrently-1.0.5.tgz",
		"integrity": "sha512-f2domd9fsVDFtaFcbaRZuYXwtdmnzqbADSwhSWYxYB/Q8zsdUUFMXVRwXGDMWmbEzAn1kdRrtI1T/KTFOL4X2A==",
		"dev": true,
		"requires": {
		  "aproba": "^1.1.1",
		  "fs-write-stream-atomic": "^1.0.8",
		  "iferr": "^0.1.5",
		  "mkdirp": "^0.5.1",
		  "rimraf": "^2.5.4",
		  "run-queue": "^1.0.0"
		}
	  },
	  "copy-descriptor": {
		"version": "0.1.1",
		"resolved": "https://npm.intra.longguikeji.com/copy-descriptor/-/copy-descriptor-0.1.1.tgz",
		"integrity": "sha1-Z29us8OZl8LuGsOpJP1hJHSPV40=",
		"dev": true
	  },
	  "core-js-compat": {
		"version": "3.6.2",
		"resolved": "https://npm.intra.longguikeji.com/core-js-compat/-/core-js-compat-3.6.2.tgz",
		"integrity": "sha512-+G28dzfYGtAM+XGvB1C5AS1ZPKfQ47HLhcdeIQdZgQnJVdp7/D0m+W/TErwhgsX6CujRUk/LebB6dCrKrtJrvQ==",
		"dev": true,
		"requires": {
		  "browserslist": "^4.8.3",
		  "semver": "7.0.0"
		},
		"dependencies": {
		  "semver": {
			"version": "7.0.0",
			"resolved": "https://npm.intra.longguikeji.com/semver/-/semver-7.0.0.tgz",
			"integrity": "sha512-+GB6zVA9LWh6zovYQLALHwv5rb2PHGlJi3lfiqIHxR0uuwCgefcOJc59v9fv1w8GbStwxuuqqAjI9NMAOOgq1A==",
			"dev": true
		  }
		}
	  },
	  "core-util-is": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/core-util-is/-/core-util-is-1.0.2.tgz",
		"integrity": "sha1-tf1UIgqivFq1eqtxQMlAdUUDwac=",
		"dev": true
	  },
	  "create-ecdh": {
		"version": "4.0.3",
		"resolved": "https://npm.intra.longguikeji.com/create-ecdh/-/create-ecdh-4.0.3.tgz",
		"integrity": "sha512-GbEHQPMOswGpKXM9kCWVrremUcBmjteUaQ01T9rkKCPDXfUHX0IoP9LpHYo2NPFampa4e+/pFDc3jQdxrxQLaw==",
		"dev": true,
		"requires": {
		  "bn.js": "^4.1.0",
		  "elliptic": "^6.0.0"
		}
	  },
	  "create-hash": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/create-hash/-/create-hash-1.2.0.tgz",
		"integrity": "sha512-z00bCGNHDG8mHAkP7CtT1qVu+bFQUPjYq/4Iv3C3kWjTFV10zIjfSoeqXo9Asws8gwSHDGj/hl2u4OGIjapeCg==",
		"dev": true,
		"requires": {
		  "cipher-base": "^1.0.1",
		  "inherits": "^2.0.1",
		  "md5.js": "^1.3.4",
		  "ripemd160": "^2.0.1",
		  "sha.js": "^2.4.0"
		}
	  },
	  "create-hmac": {
		"version": "1.1.7",
		"resolved": "https://npm.intra.longguikeji.com/create-hmac/-/create-hmac-1.1.7.tgz",
		"integrity": "sha512-MJG9liiZ+ogc4TzUwuvbER1JRdgvUFSB5+VR/g5h82fGaIRWMWddtKBHi7/sVhfjQZ6SehlyhvQYrcYkaUIpLg==",
		"dev": true,
		"requires": {
		  "cipher-base": "^1.0.3",
		  "create-hash": "^1.1.0",
		  "inherits": "^2.0.1",
		  "ripemd160": "^2.0.0",
		  "safe-buffer": "^5.0.1",
		  "sha.js": "^2.4.8"
		}
	  },
	  "cross-spawn": {
		"version": "6.0.5",
		"resolved": "https://npm.intra.longguikeji.com/cross-spawn/-/cross-spawn-6.0.5.tgz",
		"integrity": "sha512-eTVLrBSt7fjbDygz805pMnstIs2VTBNkRm0qxZd+M7A5XDdxVRWO5MxGBXZhjY4cqLYLdtrGqRf8mBPmzwSpWQ==",
		"dev": true,
		"requires": {
		  "nice-try": "^1.0.4",
		  "path-key": "^2.0.1",
		  "semver": "^5.5.0",
		  "shebang-command": "^1.2.0",
		  "which": "^1.2.9"
		}
	  },
	  "crypto-browserify": {
		"version": "3.12.0",
		"resolved": "https://npm.intra.longguikeji.com/crypto-browserify/-/crypto-browserify-3.12.0.tgz",
		"integrity": "sha512-fz4spIh+znjO2VjL+IdhEpRJ3YN6sMzITSBijk6FK2UvTqruSQW+/cCZTSNsMiZNvUeq0CqurF+dAbyiGOY6Wg==",
		"dev": true,
		"requires": {
		  "browserify-cipher": "^1.0.0",
		  "browserify-sign": "^4.0.0",
		  "create-ecdh": "^4.0.0",
		  "create-hash": "^1.1.0",
		  "create-hmac": "^1.1.0",
		  "diffie-hellman": "^5.0.0",
		  "inherits": "^2.0.1",
		  "pbkdf2": "^3.0.3",
		  "public-encrypt": "^4.0.0",
		  "randombytes": "^2.0.0",
		  "randomfill": "^1.0.3"
		}
	  },
	  "cyclist": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/cyclist/-/cyclist-1.0.1.tgz",
		"integrity": "sha1-WW6WmP0MgOEgOMK4LW6xs1tiJNk=",
		"dev": true
	  },
	  "debug": {
		"version": "4.1.1",
		"resolved": "https://npm.intra.longguikeji.com/debug/-/debug-4.1.1.tgz",
		"integrity": "sha512-pYAIzeRo8J6KPEaJ0VWOh5Pzkbw/RetuzehGM7QRRX5he4fPHx2rdKMB256ehJCkX+XRQm16eZLqLNS8RSZXZw==",
		"requires": {
		  "ms": "^2.1.1"
		},
		"dependencies": {
		  "ms": {
			"version": "2.1.2",
			"resolved": "https://npm.intra.longguikeji.com/ms/-/ms-2.1.2.tgz",
			"integrity": "sha512-sGkPx+VjMtmA6MX27oA4FBFELFCZZ4S4XqeGOXCv68tT+jb3vk/RyaKWP0PTKyWtmLSM0b+adUTEvbs1PEaH2w=="
		  }
		}
	  },
	  "decamelize": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/decamelize/-/decamelize-1.2.0.tgz",
		"integrity": "sha1-9lNNFRSCabIDUue+4m9QH5oZEpA="
	  },
	  "decode-uri-component": {
		"version": "0.2.0",
		"resolved": "https://npm.intra.longguikeji.com/decode-uri-component/-/decode-uri-component-0.2.0.tgz",
		"integrity": "sha1-6zkTMzRYd1y4TNGh+uBiEGu4dUU=",
		"dev": true
	  },
	  "define-properties": {
		"version": "1.1.3",
		"resolved": "https://npm.intra.longguikeji.com/define-properties/-/define-properties-1.1.3.tgz",
		"integrity": "sha512-3MqfYKj2lLzdMSf8ZIZE/V+Zuy+BgD6f164e8K2w7dgnpKArBDerGYpM46IYYcjnkdPNMjPk9A6VFB8+3SKlXQ==",
		"dev": true,
		"requires": {
		  "object-keys": "^1.0.12"
		}
	  },
	  "define-property": {
		"version": "2.0.2",
		"resolved": "https://npm.intra.longguikeji.com/define-property/-/define-property-2.0.2.tgz",
		"integrity": "sha512-jwK2UV4cnPpbcG7+VRARKTZPUWowwXA8bzH5NP6ud0oeAxyYPuGZUAC7hMugpCdz4BeSZl2Dl9k66CHJ/46ZYQ==",
		"dev": true,
		"requires": {
		  "is-descriptor": "^1.0.2",
		  "isobject": "^3.0.1"
		},
		"dependencies": {
		  "is-accessor-descriptor": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/is-accessor-descriptor/-/is-accessor-descriptor-1.0.0.tgz",
			"integrity": "sha512-m5hnHTkcVsPfqx3AKlyttIPb7J+XykHvJP2B9bZDjlhLIoEq4XoK64Vg7boZlVWYK6LUY94dYPEE7Lh0ZkZKcQ==",
			"dev": true,
			"requires": {
			  "kind-of": "^6.0.0"
			}
		  },
		  "is-data-descriptor": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/is-data-descriptor/-/is-data-descriptor-1.0.0.tgz",
			"integrity": "sha512-jbRXy1FmtAoCjQkVmIVYwuuqDFUbaOeDjmed1tOGPrsMhtJA4rD9tkgA0F1qJ3gRFRXcHYVkdeaP50Q5rE/jLQ==",
			"dev": true,
			"requires": {
			  "kind-of": "^6.0.0"
			}
		  },
		  "is-descriptor": {
			"version": "1.0.2",
			"resolved": "https://npm.intra.longguikeji.com/is-descriptor/-/is-descriptor-1.0.2.tgz",
			"integrity": "sha512-2eis5WqQGV7peooDyLmNEPUrps9+SXX5c9pL3xEB+4e9HnGuDa7mB7kHxHw4CbqS9k1T2hOH3miL8n8WtiYVtg==",
			"dev": true,
			"requires": {
			  "is-accessor-descriptor": "^1.0.0",
			  "is-data-descriptor": "^1.0.0",
			  "kind-of": "^6.0.2"
			}
		  }
		}
	  },
	  "depd": {
		"version": "1.1.2",
		"resolved": "https://npm.intra.longguikeji.com/depd/-/depd-1.1.2.tgz",
		"integrity": "sha1-m81S4UwJd2PnSbJ0xDRu0uVgtak="
	  },
	  "des.js": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/des.js/-/des.js-1.0.1.tgz",
		"integrity": "sha512-Q0I4pfFrv2VPd34/vfLrFOoRmlYj3OV50i7fskps1jZWK1kApMWWT9G6RRUeYedLcBDIhnSDaUvJMb3AhUlaEA==",
		"dev": true,
		"requires": {
		  "inherits": "^2.0.1",
		  "minimalistic-assert": "^1.0.0"
		}
	  },
	  "destroy": {
		"version": "1.0.4",
		"resolved": "https://npm.intra.longguikeji.com/destroy/-/destroy-1.0.4.tgz",
		"integrity": "sha1-l4hXRCxEdJ5CBmE+N5RiBYJqvYA="
	  },
	  "detect-file": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/detect-file/-/detect-file-1.0.0.tgz",
		"integrity": "sha1-8NZtA2cqglyxtzvbP+YjEMjlUrc=",
		"dev": true
	  },
	  "diffie-hellman": {
		"version": "5.0.3",
		"resolved": "https://npm.intra.longguikeji.com/diffie-hellman/-/diffie-hellman-5.0.3.tgz",
		"integrity": "sha512-kqag/Nl+f3GwyK25fhUMYj81BUOrZ9IuJsjIcDE5icNM9FJHAVm3VcUDxdLPoQtTuUylWm6ZIknYJwwaPxsUzg==",
		"dev": true,
		"requires": {
		  "bn.js": "^4.1.0",
		  "miller-rabin": "^4.0.0",
		  "randombytes": "^2.0.0"
		}
	  },
	  "domain-browser": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/domain-browser/-/domain-browser-1.2.0.tgz",
		"integrity": "sha512-jnjyiM6eRyZl2H+W8Q/zLMA481hzi0eszAaBUzIVnmYVDBbnLxVNnfu1HgEBvCbL+71FrxMl3E6lpKH7Ge3OXA==",
		"dev": true
	  },
	  "duplexify": {
		"version": "3.7.1",
		"resolved": "https://npm.intra.longguikeji.com/duplexify/-/duplexify-3.7.1.tgz",
		"integrity": "sha512-07z8uv2wMyS51kKhD1KsdXJg5WQ6t93RneqRxUHnskXVtlYYkLqM0gqStQZ3pj073g687jPCHrqNfCzawLYh5g==",
		"dev": true,
		"requires": {
		  "end-of-stream": "^1.0.0",
		  "inherits": "^2.0.1",
		  "readable-stream": "^2.0.0",
		  "stream-shift": "^1.0.0"
		}
	  },
	  "ee-first": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/ee-first/-/ee-first-1.1.1.tgz",
		"integrity": "sha1-WQxhFWsK4vTwJVcyoViyZrxWsh0="
	  },
	  "electron-to-chromium": {
		"version": "1.3.328",
		"resolved": "https://npm.intra.longguikeji.com/electron-to-chromium/-/electron-to-chromium-1.3.328.tgz",
		"integrity": "sha512-x4XefnFxDxFwaQ01d/pppJP9meWhOIJ/gtI6/4jqkpsadq79uL7NYSaX64naLmJqvzUBjSrO3IM2+1b/W9KdPg==",
		"dev": true
	  },
	  "elliptic": {
		"version": "6.5.2",
		"resolved": "https://npm.intra.longguikeji.com/elliptic/-/elliptic-6.5.2.tgz",
		"integrity": "sha512-f4x70okzZbIQl/NSRLkI/+tteV/9WqL98zx+SQ69KbXxmVrmjwsNUPn/gYJJ0sHvEak24cZgHIPegRePAtA/xw==",
		"dev": true,
		"requires": {
		  "bn.js": "^4.4.0",
		  "brorand": "^1.0.1",
		  "hash.js": "^1.0.0",
		  "hmac-drbg": "^1.0.0",
		  "inherits": "^2.0.1",
		  "minimalistic-assert": "^1.0.0",
		  "minimalistic-crypto-utils": "^1.0.0"
		}
	  },
	  "emoji-regex": {
		"version": "8.0.0",
		"resolved": "https://npm.intra.longguikeji.com/emoji-regex/-/emoji-regex-8.0.0.tgz",
		"integrity": "sha512-MSjYzcWNOA0ewAHpz0MxpYFvwg6yjy1NG3xteoqz644VCo/RPgnr1/GGt+ic3iJTzQ8Eu3TdM14SawnVUmGE6A=="
	  },
	  "emojis-list": {
		"version": "2.1.0",
		"resolved": "https://npm.intra.longguikeji.com/emojis-list/-/emojis-list-2.1.0.tgz",
		"integrity": "sha1-TapNnbAPmBmIDHn6RXrlsJof04k=",
		"dev": true
	  },
	  "encodeurl": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/encodeurl/-/encodeurl-1.0.2.tgz",
		"integrity": "sha1-rT/0yG7C0CkyL1oCw6mmBslbP1k="
	  },
	  "end-of-stream": {
		"version": "1.4.4",
		"resolved": "https://npm.intra.longguikeji.com/end-of-stream/-/end-of-stream-1.4.4.tgz",
		"integrity": "sha512-+uw1inIHVPQoaVuHzRyXd21icM+cnt4CzD5rW+NC1wjOUSTOs+Te7FOv7AhN7vS9x/oIyhLP5PR1H+phQAHu5Q==",
		"dev": true,
		"requires": {
		  "once": "^1.4.0"
		}
	  },
	  "enhanced-resolve": {
		"version": "4.1.1",
		"resolved": "https://npm.intra.longguikeji.com/enhanced-resolve/-/enhanced-resolve-4.1.1.tgz",
		"integrity": "sha512-98p2zE+rL7/g/DzMHMTF4zZlCgeVdJ7yr6xzEpJRYwFYrGi9ANdn5DnJURg6RpBkyk60XYDnWIv51VfIhfNGuA==",
		"dev": true,
		"requires": {
		  "graceful-fs": "^4.1.2",
		  "memory-fs": "^0.5.0",
		  "tapable": "^1.0.0"
		},
		"dependencies": {
		  "memory-fs": {
			"version": "0.5.0",
			"resolved": "https://npm.intra.longguikeji.com/memory-fs/-/memory-fs-0.5.0.tgz",
			"integrity": "sha512-jA0rdU5KoQMC0e6ppoNRtpp6vjFq6+NY7r8hywnC7V+1Xj/MtHwGIbB1QaK/dunyjWteJzmkpd7ooeWg10T7GA==",
			"dev": true,
			"requires": {
			  "errno": "^0.1.3",
			  "readable-stream": "^2.0.1"
			}
		  }
		}
	  },
	  "errno": {
		"version": "0.1.7",
		"resolved": "https://npm.intra.longguikeji.com/errno/-/errno-0.1.7.tgz",
		"integrity": "sha512-MfrRBDWzIWifgq6tJj60gkAwtLNb6sQPlcFrSOflcP1aFmmruKQ2wRnze/8V6kgyz7H3FF8Npzv78mZ7XLLflg==",
		"dev": true,
		"requires": {
		  "prr": "~1.0.1"
		}
	  },
	  "escape-html": {
		"version": "1.0.3",
		"resolved": "https://npm.intra.longguikeji.com/escape-html/-/escape-html-1.0.3.tgz",
		"integrity": "sha1-Aljq5NPQwJdN4cFpGI7wBR0dGYg="
	  },
	  "escape-string-regexp": {
		"version": "1.0.5",
		"resolved": "https://npm.intra.longguikeji.com/escape-string-regexp/-/escape-string-regexp-1.0.5.tgz",
		"integrity": "sha1-G2HAViGQqN/2rjuyzwIAyhMLhtQ=",
		"dev": true
	  },
	  "eslint-scope": {
		"version": "4.0.3",
		"resolved": "https://npm.intra.longguikeji.com/eslint-scope/-/eslint-scope-4.0.3.tgz",
		"integrity": "sha512-p7VutNr1O/QrxysMo3E45FjYDTeXBy0iTltPFNSqKAIfjDSXC+4dj+qfyuD8bfAXrW/y6lW3O76VaYNPKfpKrg==",
		"dev": true,
		"requires": {
		  "esrecurse": "^4.1.0",
		  "estraverse": "^4.1.1"
		}
	  },
	  "esrecurse": {
		"version": "4.2.1",
		"resolved": "https://npm.intra.longguikeji.com/esrecurse/-/esrecurse-4.2.1.tgz",
		"integrity": "sha512-64RBB++fIOAXPw3P9cy89qfMlvZEXZkqqJkjqqXIvzP5ezRZjW+lPWjw35UX/3EhUPFYbg5ER4JYgDw4007/DQ==",
		"dev": true,
		"requires": {
		  "estraverse": "^4.1.0"
		}
	  },
	  "estraverse": {
		"version": "4.3.0",
		"resolved": "https://npm.intra.longguikeji.com/estraverse/-/estraverse-4.3.0.tgz",
		"integrity": "sha512-39nnKffWz8xN1BU/2c79n9nB9HDzo0niYUqx6xyqUnyoAnQyyWpOTdZEeiCch8BBu515t4wp9ZmgVfVhn9EBpw==",
		"dev": true
	  },
	  "esutils": {
		"version": "2.0.3",
		"resolved": "https://npm.intra.longguikeji.com/esutils/-/esutils-2.0.3.tgz",
		"integrity": "sha512-kVscqXk4OCp68SZ0dkgEKVi6/8ij300KBWTJq32P/dYeWTSwK41WyTxalN1eRmA5Z9UU/LX9D7FWSmV9SAYx6g==",
		"dev": true
	  },
	  "etag": {
		"version": "1.8.1",
		"resolved": "https://npm.intra.longguikeji.com/etag/-/etag-1.8.1.tgz",
		"integrity": "sha1-Qa4u62XvpiJorr/qg6x9eSmbCIc="
	  },
	  "events": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/events/-/events-3.0.0.tgz",
		"integrity": "sha512-Dc381HFWJzEOhQ+d8pkNon++bk9h6cdAoAj4iE6Q4y6xgTzySWXlKn05/TVNpjnfRqi/X0EpJEJohPjNI3zpVA==",
		"dev": true
	  },
	  "evp_bytestokey": {
		"version": "1.0.3",
		"resolved": "https://npm.intra.longguikeji.com/evp_bytestokey/-/evp_bytestokey-1.0.3.tgz",
		"integrity": "sha512-/f2Go4TognH/KvCISP7OUsHn85hT9nUkxxA9BEWxFn+Oj9o8ZNLm/40hdlgSLyuOimsrTKLUMEorQexp/aPQeA==",
		"dev": true,
		"requires": {
		  "md5.js": "^1.3.4",
		  "safe-buffer": "^5.1.1"
		}
	  },
	  "execa": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/execa/-/execa-1.0.0.tgz",
		"integrity": "sha512-adbxcyWV46qiHyvSp50TKt05tB4tK3HcmF7/nxfAdhnox83seTDbwnaqKO4sXRy7roHAIFqJP/Rw/AuEbX61LA==",
		"dev": true,
		"requires": {
		  "cross-spawn": "^6.0.0",
		  "get-stream": "^4.0.0",
		  "is-stream": "^1.1.0",
		  "npm-run-path": "^2.0.0",
		  "p-finally": "^1.0.0",
		  "signal-exit": "^3.0.0",
		  "strip-eof": "^1.0.0"
		}
	  },
	  "expand-brackets": {
		"version": "2.1.4",
		"resolved": "https://npm.intra.longguikeji.com/expand-brackets/-/expand-brackets-2.1.4.tgz",
		"integrity": "sha1-t3c14xXOMPa27/D4OwQVGiJEliI=",
		"dev": true,
		"requires": {
		  "debug": "^2.3.3",
		  "define-property": "^0.2.5",
		  "extend-shallow": "^2.0.1",
		  "posix-character-classes": "^0.1.0",
		  "regex-not": "^1.0.0",
		  "snapdragon": "^0.8.1",
		  "to-regex": "^3.0.1"
		},
		"dependencies": {
		  "debug": {
			"version": "2.6.9",
			"resolved": "https://npm.intra.longguikeji.com/debug/-/debug-2.6.9.tgz",
			"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
			"dev": true,
			"requires": {
			  "ms": "2.0.0"
			}
		  },
		  "define-property": {
			"version": "0.2.5",
			"resolved": "https://npm.intra.longguikeji.com/define-property/-/define-property-0.2.5.tgz",
			"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
			"dev": true,
			"requires": {
			  "is-descriptor": "^0.1.0"
			}
		  },
		  "extend-shallow": {
			"version": "2.0.1",
			"resolved": "https://npm.intra.longguikeji.com/extend-shallow/-/extend-shallow-2.0.1.tgz",
			"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
			"dev": true,
			"requires": {
			  "is-extendable": "^0.1.0"
			}
		  }
		}
	  },
	  "expand-tilde": {
		"version": "2.0.2",
		"resolved": "https://npm.intra.longguikeji.com/expand-tilde/-/expand-tilde-2.0.2.tgz",
		"integrity": "sha1-l+gBqgUt8CRU3kawK/YhZCzchQI=",
		"dev": true,
		"requires": {
		  "homedir-polyfill": "^1.0.1"
		}
	  },
	  "express": {
		"version": "4.17.1",
		"resolved": "https://npm.intra.longguikeji.com/express/-/express-4.17.1.tgz",
		"integrity": "sha512-mHJ9O79RqluphRrcw2X/GTh3k9tVv8YcoyY4Kkh4WDMUYKRZUq0h1o0w2rrrxBqM7VoeUVqgb27xlEMXTnYt4g==",
		"requires": {
		  "accepts": "~1.3.7",
		  "array-flatten": "1.1.1",
		  "body-parser": "1.19.0",
		  "content-disposition": "0.5.3",
		  "content-type": "~1.0.4",
		  "cookie": "0.4.0",
		  "cookie-signature": "1.0.6",
		  "debug": "2.6.9",
		  "depd": "~1.1.2",
		  "encodeurl": "~1.0.2",
		  "escape-html": "~1.0.3",
		  "etag": "~1.8.1",
		  "finalhandler": "~1.1.2",
		  "fresh": "0.5.2",
		  "merge-descriptors": "1.0.1",
		  "methods": "~1.1.2",
		  "on-finished": "~2.3.0",
		  "parseurl": "~1.3.3",
		  "path-to-regexp": "0.1.7",
		  "proxy-addr": "~2.0.5",
		  "qs": "6.7.0",
		  "range-parser": "~1.2.1",
		  "safe-buffer": "5.1.2",
		  "send": "0.17.1",
		  "serve-static": "1.14.1",
		  "setprototypeof": "1.1.1",
		  "statuses": "~1.5.0",
		  "type-is": "~1.6.18",
		  "utils-merge": "1.0.1",
		  "vary": "~1.1.2"
		},
		"dependencies": {
		  "debug": {
			"version": "2.6.9",
			"resolved": "https://npm.intra.longguikeji.com/debug/-/debug-2.6.9.tgz",
			"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
			"requires": {
			  "ms": "2.0.0"
			}
		  }
		}
	  },
	  "express-formidable": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/express-formidable/-/express-formidable-1.2.0.tgz",
		"integrity": "sha512-w1vXjF3gb50UKTNkFaW8/4rqY4dUrKfZ1sAZzwAF9YxCAgj/29QZsycf71di0GkskrZOAkubk9pvGYfxyAMYiw==",
		"requires": {
		  "formidable": "^1.0.17"
		}
	  },
	  "extend-shallow": {
		"version": "3.0.2",
		"resolved": "https://npm.intra.longguikeji.com/extend-shallow/-/extend-shallow-3.0.2.tgz",
		"integrity": "sha1-Jqcarwc7OfshJxcnRhMcJwQCjbg=",
		"dev": true,
		"requires": {
		  "assign-symbols": "^1.0.0",
		  "is-extendable": "^1.0.1"
		},
		"dependencies": {
		  "is-extendable": {
			"version": "1.0.1",
			"resolved": "https://npm.intra.longguikeji.com/is-extendable/-/is-extendable-1.0.1.tgz",
			"integrity": "sha512-arnXMxT1hhoKo9k1LZdmlNyJdDDfy2v0fXjFlmok4+i8ul/6WlbVge9bhM74OpNPQPMGUToDtz+KXa1PneJxOA==",
			"dev": true,
			"requires": {
			  "is-plain-object": "^2.0.4"
			}
		  }
		}
	  },
	  "extglob": {
		"version": "2.0.4",
		"resolved": "https://npm.intra.longguikeji.com/extglob/-/extglob-2.0.4.tgz",
		"integrity": "sha512-Nmb6QXkELsuBr24CJSkilo6UHHgbekK5UiZgfE6UHD3Eb27YC6oD+bhcT+tJ6cl8dmsgdQxnWlcry8ksBIBLpw==",
		"dev": true,
		"requires": {
		  "array-unique": "^0.3.2",
		  "define-property": "^1.0.0",
		  "expand-brackets": "^2.1.4",
		  "extend-shallow": "^2.0.1",
		  "fragment-cache": "^0.2.1",
		  "regex-not": "^1.0.0",
		  "snapdragon": "^0.8.1",
		  "to-regex": "^3.0.1"
		},
		"dependencies": {
		  "define-property": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/define-property/-/define-property-1.0.0.tgz",
			"integrity": "sha1-dp66rz9KY6rTr56NMEybvnm/sOY=",
			"dev": true,
			"requires": {
			  "is-descriptor": "^1.0.0"
			}
		  },
		  "extend-shallow": {
			"version": "2.0.1",
			"resolved": "https://npm.intra.longguikeji.com/extend-shallow/-/extend-shallow-2.0.1.tgz",
			"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
			"dev": true,
			"requires": {
			  "is-extendable": "^0.1.0"
			}
		  },
		  "is-accessor-descriptor": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/is-accessor-descriptor/-/is-accessor-descriptor-1.0.0.tgz",
			"integrity": "sha512-m5hnHTkcVsPfqx3AKlyttIPb7J+XykHvJP2B9bZDjlhLIoEq4XoK64Vg7boZlVWYK6LUY94dYPEE7Lh0ZkZKcQ==",
			"dev": true,
			"requires": {
			  "kind-of": "^6.0.0"
			}
		  },
		  "is-data-descriptor": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/is-data-descriptor/-/is-data-descriptor-1.0.0.tgz",
			"integrity": "sha512-jbRXy1FmtAoCjQkVmIVYwuuqDFUbaOeDjmed1tOGPrsMhtJA4rD9tkgA0F1qJ3gRFRXcHYVkdeaP50Q5rE/jLQ==",
			"dev": true,
			"requires": {
			  "kind-of": "^6.0.0"
			}
		  },
		  "is-descriptor": {
			"version": "1.0.2",
			"resolved": "https://npm.intra.longguikeji.com/is-descriptor/-/is-descriptor-1.0.2.tgz",
			"integrity": "sha512-2eis5WqQGV7peooDyLmNEPUrps9+SXX5c9pL3xEB+4e9HnGuDa7mB7kHxHw4CbqS9k1T2hOH3miL8n8WtiYVtg==",
			"dev": true,
			"requires": {
			  "is-accessor-descriptor": "^1.0.0",
			  "is-data-descriptor": "^1.0.0",
			  "kind-of": "^6.0.2"
			}
		  }
		}
	  },
	  "fast-deep-equal": {
		"version": "2.0.1",
		"resolved": "https://npm.intra.longguikeji.com/fast-deep-equal/-/fast-deep-equal-2.0.1.tgz",
		"integrity": "sha1-ewUhjd+WZ79/Nwv3/bLLFf3Qqkk=",
		"dev": true
	  },
	  "fast-json-stable-stringify": {
		"version": "2.1.0",
		"resolved": "https://npm.intra.longguikeji.com/fast-json-stable-stringify/-/fast-json-stable-stringify-2.1.0.tgz",
		"integrity": "sha512-lhd/wF+Lk98HZoTCtlVraHtfh5XYijIjalXck7saUtuanSDyLMxnHhSXEDJqHxD7msR8D0uCmqlkwjCV8xvwHw==",
		"dev": true
	  },
	  "figgy-pudding": {
		"version": "3.5.1",
		"resolved": "https://npm.intra.longguikeji.com/figgy-pudding/-/figgy-pudding-3.5.1.tgz",
		"integrity": "sha512-vNKxJHTEKNThjfrdJwHc7brvM6eVevuO5nTj6ez8ZQ1qbXTvGthucRF7S4vf2cr71QVnT70V34v0S1DyQsti0w==",
		"dev": true
	  },
	  "file-uri-to-path": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/file-uri-to-path/-/file-uri-to-path-1.0.0.tgz",
		"integrity": "sha512-0Zt+s3L7Vf1biwWZ29aARiVYLx7iMGnEUl9x33fbB/j3jR81u/O2LbqK+Bm1CDSNDKVtJ/YjwY7TUd5SkeLQLw==",
		"dev": true,
		"optional": true
	  },
	  "fill-range": {
		"version": "4.0.0",
		"resolved": "https://npm.intra.longguikeji.com/fill-range/-/fill-range-4.0.0.tgz",
		"integrity": "sha1-1USBHUKPmOsGpj3EAtJAPDKMOPc=",
		"dev": true,
		"requires": {
		  "extend-shallow": "^2.0.1",
		  "is-number": "^3.0.0",
		  "repeat-string": "^1.6.1",
		  "to-regex-range": "^2.1.0"
		},
		"dependencies": {
		  "extend-shallow": {
			"version": "2.0.1",
			"resolved": "https://npm.intra.longguikeji.com/extend-shallow/-/extend-shallow-2.0.1.tgz",
			"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
			"dev": true,
			"requires": {
			  "is-extendable": "^0.1.0"
			}
		  }
		}
	  },
	  "finalhandler": {
		"version": "1.1.2",
		"resolved": "https://npm.intra.longguikeji.com/finalhandler/-/finalhandler-1.1.2.tgz",
		"integrity": "sha512-aAWcW57uxVNrQZqFXjITpW3sIUQmHGG3qSb9mUah9MgMC4NeWhNOlNjXEYq3HjRAvL6arUviZGGJsBg6z0zsWA==",
		"requires": {
		  "debug": "2.6.9",
		  "encodeurl": "~1.0.2",
		  "escape-html": "~1.0.3",
		  "on-finished": "~2.3.0",
		  "parseurl": "~1.3.3",
		  "statuses": "~1.5.0",
		  "unpipe": "~1.0.0"
		},
		"dependencies": {
		  "debug": {
			"version": "2.6.9",
			"resolved": "https://npm.intra.longguikeji.com/debug/-/debug-2.6.9.tgz",
			"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
			"requires": {
			  "ms": "2.0.0"
			}
		  }
		}
	  },
	  "find-cache-dir": {
		"version": "2.1.0",
		"resolved": "https://npm.intra.longguikeji.com/find-cache-dir/-/find-cache-dir-2.1.0.tgz",
		"integrity": "sha512-Tq6PixE0w/VMFfCgbONnkiQIVol/JJL7nRMi20fqzA4NRs9AfeqMGeRdPi3wIhYkxjeBaWh2rxwapn5Tu3IqOQ==",
		"dev": true,
		"requires": {
		  "commondir": "^1.0.1",
		  "make-dir": "^2.0.0",
		  "pkg-dir": "^3.0.0"
		}
	  },
	  "find-up": {
		"version": "4.1.0",
		"resolved": "https://npm.intra.longguikeji.com/find-up/-/find-up-4.1.0.tgz",
		"integrity": "sha512-PpOwAdQ/YlXQ2vj8a3h8IipDuYRi3wceVQQGYWxNINccq40Anw7BlsEXCMbt1Zt+OLA6Fq9suIpIWD0OsnISlw==",
		"requires": {
		  "locate-path": "^5.0.0",
		  "path-exists": "^4.0.0"
		}
	  },
	  "findup-sync": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/findup-sync/-/findup-sync-3.0.0.tgz",
		"integrity": "sha512-YbffarhcicEhOrm4CtrwdKBdCuz576RLdhJDsIfvNtxUuhdRet1qZcsMjqbePtAseKdAnDyM/IyXbu7PRPRLYg==",
		"dev": true,
		"requires": {
		  "detect-file": "^1.0.0",
		  "is-glob": "^4.0.0",
		  "micromatch": "^3.0.4",
		  "resolve-dir": "^1.0.1"
		}
	  },
	  "flush-write-stream": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/flush-write-stream/-/flush-write-stream-1.1.1.tgz",
		"integrity": "sha512-3Z4XhFZ3992uIq0XOqb9AreonueSYphE6oYbpt5+3u06JWklbsPkNv3ZKkP9Bz/r+1MWCaMoSQ28P85+1Yc77w==",
		"dev": true,
		"requires": {
		  "inherits": "^2.0.3",
		  "readable-stream": "^2.3.6"
		}
	  },
	  "follow-redirects": {
		"version": "1.5.10",
		"resolved": "https://npm.intra.longguikeji.com/follow-redirects/-/follow-redirects-1.5.10.tgz",
		"integrity": "sha512-0V5l4Cizzvqt5D44aTXbFZz+FtyXV1vrDN6qrelxtfYQKW0KO0W2T/hkE8xvGa/540LkZlkaUjO4ailYTFtHVQ==",
		"requires": {
		  "debug": "=3.1.0"
		},
		"dependencies": {
		  "debug": {
			"version": "3.1.0",
			"resolved": "https://npm.intra.longguikeji.com/debug/-/debug-3.1.0.tgz",
			"integrity": "sha512-OX8XqP7/1a9cqkxYw2yXss15f26NKWBpDXQd0/uK/KPqdQhxbPa994hnzjcE2VqQpDslf55723cKPUOGSmMY3g==",
			"requires": {
			  "ms": "2.0.0"
			}
		  }
		}
	  },
	  "for-in": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/for-in/-/for-in-1.0.2.tgz",
		"integrity": "sha1-gQaNKVqBQuwKxybG4iAMMPttXoA=",
		"dev": true
	  },
	  "formidable": {
		"version": "1.2.1",
		"resolved": "https://npm.intra.longguikeji.com/formidable/-/formidable-1.2.1.tgz",
		"integrity": "sha512-Fs9VRguL0gqGHkXS5GQiMCr1VhZBxz0JnJs4JmMp/2jL18Fmbzvv7vOFRU+U8TBkHEE/CX1qDXzJplVULgsLeg=="
	  },
	  "forwarded": {
		"version": "0.1.2",
		"resolved": "https://npm.intra.longguikeji.com/forwarded/-/forwarded-0.1.2.tgz",
		"integrity": "sha1-mMI9qxF1ZXuMBXPozszZGw/xjIQ="
	  },
	  "fragment-cache": {
		"version": "0.2.1",
		"resolved": "https://npm.intra.longguikeji.com/fragment-cache/-/fragment-cache-0.2.1.tgz",
		"integrity": "sha1-QpD60n8T6Jvn8zeZxrxaCr//DRk=",
		"dev": true,
		"requires": {
		  "map-cache": "^0.2.2"
		}
	  },
	  "fresh": {
		"version": "0.5.2",
		"resolved": "https://npm.intra.longguikeji.com/fresh/-/fresh-0.5.2.tgz",
		"integrity": "sha1-PYyt2Q2XZWn6g1qx+OSyOhBWBac="
	  },
	  "from2": {
		"version": "2.3.0",
		"resolved": "https://npm.intra.longguikeji.com/from2/-/from2-2.3.0.tgz",
		"integrity": "sha1-i/tVAr3kpNNs/e6gB/zKIdfjgq8=",
		"dev": true,
		"requires": {
		  "inherits": "^2.0.1",
		  "readable-stream": "^2.0.0"
		}
	  },
	  "fs-readdir-recursive": {
		"version": "1.1.0",
		"resolved": "https://npm.intra.longguikeji.com/fs-readdir-recursive/-/fs-readdir-recursive-1.1.0.tgz",
		"integrity": "sha512-GNanXlVr2pf02+sPN40XN8HG+ePaNcvM0q5mZBd668Obwb0yD5GiUbZOFgwn8kGMY6I3mdyDJzieUy3PTYyTRA==",
		"dev": true
	  },
	  "fs-write-stream-atomic": {
		"version": "1.0.10",
		"resolved": "https://npm.intra.longguikeji.com/fs-write-stream-atomic/-/fs-write-stream-atomic-1.0.10.tgz",
		"integrity": "sha1-tH31NJPvkR33VzHnCp3tAYnbQMk=",
		"dev": true,
		"requires": {
		  "graceful-fs": "^4.1.2",
		  "iferr": "^0.1.5",
		  "imurmurhash": "^0.1.4",
		  "readable-stream": "1 || 2"
		}
	  },
	  "fs.realpath": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/fs.realpath/-/fs.realpath-1.0.0.tgz",
		"integrity": "sha1-FQStJSMVjKpA20onh8sBQRmU6k8=",
		"dev": true
	  },
	  "fsevents": {
		"version": "1.2.11",
		"resolved": "https://npm.intra.longguikeji.com/fsevents/-/fsevents-1.2.11.tgz",
		"integrity": "sha512-+ux3lx6peh0BpvY0JebGyZoiR4D+oYzdPZMKJwkZ+sFkNJzpL7tXc/wehS49gUAxg3tmMHPHZkA8JU2rhhgDHw==",
		"dev": true,
		"optional": true,
		"requires": {
		  "bindings": "^1.5.0",
		  "nan": "^2.12.1",
		  "node-pre-gyp": "*"
		},
		"dependencies": {
		  "abbrev": {
			"version": "1.1.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "ansi-regex": {
			"version": "2.1.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "aproba": {
			"version": "1.2.0",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "are-we-there-yet": {
			"version": "1.1.5",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "delegates": "^1.0.0",
			  "readable-stream": "^2.0.6"
			}
		  },
		  "balanced-match": {
			"version": "1.0.0",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "brace-expansion": {
			"version": "1.1.11",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "balanced-match": "^1.0.0",
			  "concat-map": "0.0.1"
			}
		  },
		  "chownr": {
			"version": "1.1.3",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "code-point-at": {
			"version": "1.1.0",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "concat-map": {
			"version": "0.0.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "console-control-strings": {
			"version": "1.1.0",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "core-util-is": {
			"version": "1.0.2",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "debug": {
			"version": "3.2.6",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "ms": "^2.1.1"
			}
		  },
		  "deep-extend": {
			"version": "0.6.0",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "delegates": {
			"version": "1.0.0",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "detect-libc": {
			"version": "1.0.3",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "fs-minipass": {
			"version": "1.2.7",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "minipass": "^2.6.0"
			}
		  },
		  "fs.realpath": {
			"version": "1.0.0",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "gauge": {
			"version": "2.7.4",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "aproba": "^1.0.3",
			  "console-control-strings": "^1.0.0",
			  "has-unicode": "^2.0.0",
			  "object-assign": "^4.1.0",
			  "signal-exit": "^3.0.0",
			  "string-width": "^1.0.1",
			  "strip-ansi": "^3.0.1",
			  "wide-align": "^1.1.0"
			}
		  },
		  "glob": {
			"version": "7.1.6",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "fs.realpath": "^1.0.0",
			  "inflight": "^1.0.4",
			  "inherits": "2",
			  "minimatch": "^3.0.4",
			  "once": "^1.3.0",
			  "path-is-absolute": "^1.0.0"
			}
		  },
		  "has-unicode": {
			"version": "2.0.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "iconv-lite": {
			"version": "0.4.24",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "safer-buffer": ">= 2.1.2 < 3"
			}
		  },
		  "ignore-walk": {
			"version": "3.0.3",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "minimatch": "^3.0.4"
			}
		  },
		  "inflight": {
			"version": "1.0.6",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "once": "^1.3.0",
			  "wrappy": "1"
			}
		  },
		  "inherits": {
			"version": "2.0.4",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "ini": {
			"version": "1.3.5",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "is-fullwidth-code-point": {
			"version": "1.0.0",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "number-is-nan": "^1.0.0"
			}
		  },
		  "isarray": {
			"version": "1.0.0",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "minimatch": {
			"version": "3.0.4",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "brace-expansion": "^1.1.7"
			}
		  },
		  "minimist": {
			"version": "0.0.8",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "minipass": {
			"version": "2.9.0",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "safe-buffer": "^5.1.2",
			  "yallist": "^3.0.0"
			}
		  },
		  "minizlib": {
			"version": "1.3.3",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "minipass": "^2.9.0"
			}
		  },
		  "mkdirp": {
			"version": "0.5.1",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "minimist": "0.0.8"
			}
		  },
		  "ms": {
			"version": "2.1.2",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "needle": {
			"version": "2.4.0",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "debug": "^3.2.6",
			  "iconv-lite": "^0.4.4",
			  "sax": "^1.2.4"
			}
		  },
		  "node-pre-gyp": {
			"version": "0.14.0",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "detect-libc": "^1.0.2",
			  "mkdirp": "^0.5.1",
			  "needle": "^2.2.1",
			  "nopt": "^4.0.1",
			  "npm-packlist": "^1.1.6",
			  "npmlog": "^4.0.2",
			  "rc": "^1.2.7",
			  "rimraf": "^2.6.1",
			  "semver": "^5.3.0",
			  "tar": "^4.4.2"
			}
		  },
		  "nopt": {
			"version": "4.0.1",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "abbrev": "1",
			  "osenv": "^0.1.4"
			}
		  },
		  "npm-bundled": {
			"version": "1.1.1",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "npm-normalize-package-bin": "^1.0.1"
			}
		  },
		  "npm-normalize-package-bin": {
			"version": "1.0.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "npm-packlist": {
			"version": "1.4.7",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "ignore-walk": "^3.0.1",
			  "npm-bundled": "^1.0.1"
			}
		  },
		  "npmlog": {
			"version": "4.1.2",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "are-we-there-yet": "~1.1.2",
			  "console-control-strings": "~1.1.0",
			  "gauge": "~2.7.3",
			  "set-blocking": "~2.0.0"
			}
		  },
		  "number-is-nan": {
			"version": "1.0.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "object-assign": {
			"version": "4.1.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "once": {
			"version": "1.4.0",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "wrappy": "1"
			}
		  },
		  "os-homedir": {
			"version": "1.0.2",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "os-tmpdir": {
			"version": "1.0.2",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "osenv": {
			"version": "0.1.5",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "os-homedir": "^1.0.0",
			  "os-tmpdir": "^1.0.0"
			}
		  },
		  "path-is-absolute": {
			"version": "1.0.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "process-nextick-args": {
			"version": "2.0.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "rc": {
			"version": "1.2.8",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "deep-extend": "^0.6.0",
			  "ini": "~1.3.0",
			  "minimist": "^1.2.0",
			  "strip-json-comments": "~2.0.1"
			},
			"dependencies": {
			  "minimist": {
				"version": "1.2.0",
				"bundled": true,
				"dev": true,
				"optional": true
			  }
			}
		  },
		  "readable-stream": {
			"version": "2.3.6",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "core-util-is": "~1.0.0",
			  "inherits": "~2.0.3",
			  "isarray": "~1.0.0",
			  "process-nextick-args": "~2.0.0",
			  "safe-buffer": "~5.1.1",
			  "string_decoder": "~1.1.1",
			  "util-deprecate": "~1.0.1"
			}
		  },
		  "rimraf": {
			"version": "2.7.1",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "glob": "^7.1.3"
			}
		  },
		  "safe-buffer": {
			"version": "5.1.2",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "safer-buffer": {
			"version": "2.1.2",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "sax": {
			"version": "1.2.4",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "semver": {
			"version": "5.7.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "set-blocking": {
			"version": "2.0.0",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "signal-exit": {
			"version": "3.0.2",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "string-width": {
			"version": "1.0.2",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "code-point-at": "^1.0.0",
			  "is-fullwidth-code-point": "^1.0.0",
			  "strip-ansi": "^3.0.0"
			}
		  },
		  "string_decoder": {
			"version": "1.1.1",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "safe-buffer": "~5.1.0"
			}
		  },
		  "strip-ansi": {
			"version": "3.0.1",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "ansi-regex": "^2.0.0"
			}
		  },
		  "strip-json-comments": {
			"version": "2.0.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "tar": {
			"version": "4.4.13",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "chownr": "^1.1.1",
			  "fs-minipass": "^1.2.5",
			  "minipass": "^2.8.6",
			  "minizlib": "^1.2.1",
			  "mkdirp": "^0.5.0",
			  "safe-buffer": "^5.1.2",
			  "yallist": "^3.0.3"
			}
		  },
		  "util-deprecate": {
			"version": "1.0.2",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "wide-align": {
			"version": "1.1.3",
			"bundled": true,
			"dev": true,
			"optional": true,
			"requires": {
			  "string-width": "^1.0.2 || 2"
			}
		  },
		  "wrappy": {
			"version": "1.0.2",
			"bundled": true,
			"dev": true,
			"optional": true
		  },
		  "yallist": {
			"version": "3.1.1",
			"bundled": true,
			"dev": true,
			"optional": true
		  }
		}
	  },
	  "function-bind": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/function-bind/-/function-bind-1.1.1.tgz",
		"integrity": "sha512-yIovAzMX49sF8Yl58fSCWJ5svSLuaibPxXQJFLmBObTuCr0Mf1KiPopGM9NiFjiYBCbfaa2Fh6breQ6ANVTI0A==",
		"dev": true
	  },
	  "get-caller-file": {
		"version": "2.0.5",
		"resolved": "https://npm.intra.longguikeji.com/get-caller-file/-/get-caller-file-2.0.5.tgz",
		"integrity": "sha512-DyFP3BM/3YHTQOCUL/w0OZHR0lpKeGrxotcHWcqNEdnltqFwXVfhEBQ94eIo34AfQpo0rGki4cyIiftY06h2Fg=="
	  },
	  "get-stream": {
		"version": "4.1.0",
		"resolved": "https://npm.intra.longguikeji.com/get-stream/-/get-stream-4.1.0.tgz",
		"integrity": "sha512-GMat4EJ5161kIy2HevLlr4luNjBgvmj413KaQA7jt4V8B4RDsfpHk7WQ9GVqfYyyx8OS/L66Kox+rJRNklLK7w==",
		"dev": true,
		"requires": {
		  "pump": "^3.0.0"
		}
	  },
	  "get-value": {
		"version": "2.0.6",
		"resolved": "https://npm.intra.longguikeji.com/get-value/-/get-value-2.0.6.tgz",
		"integrity": "sha1-3BXKHGcjh8p2vTesCjlbogQqLCg=",
		"dev": true
	  },
	  "glob": {
		"version": "7.1.6",
		"resolved": "https://npm.intra.longguikeji.com/glob/-/glob-7.1.6.tgz",
		"integrity": "sha512-LwaxwyZ72Lk7vZINtNNrywX0ZuLyStrdDtabefZKAY5ZGJhVtgdznluResxNmPitE0SAO+O26sWTHeKSI2wMBA==",
		"dev": true,
		"requires": {
		  "fs.realpath": "^1.0.0",
		  "inflight": "^1.0.4",
		  "inherits": "2",
		  "minimatch": "^3.0.4",
		  "once": "^1.3.0",
		  "path-is-absolute": "^1.0.0"
		}
	  },
	  "glob-parent": {
		"version": "3.1.0",
		"resolved": "https://npm.intra.longguikeji.com/glob-parent/-/glob-parent-3.1.0.tgz",
		"integrity": "sha1-nmr2KZ2NO9K9QEMIMr0RPfkGxa4=",
		"dev": true,
		"requires": {
		  "is-glob": "^3.1.0",
		  "path-dirname": "^1.0.0"
		},
		"dependencies": {
		  "is-glob": {
			"version": "3.1.0",
			"resolved": "https://npm.intra.longguikeji.com/is-glob/-/is-glob-3.1.0.tgz",
			"integrity": "sha1-e6WuJCF4BKxwcHuWkiVnSGzD6Eo=",
			"dev": true,
			"requires": {
			  "is-extglob": "^2.1.0"
			}
		  }
		}
	  },
	  "global-modules": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/global-modules/-/global-modules-2.0.0.tgz",
		"integrity": "sha512-NGbfmJBp9x8IxyJSd1P+otYK8vonoJactOogrVfFRIAEY1ukil8RSKDz2Yo7wh1oihl51l/r6W4epkeKJHqL8A==",
		"dev": true,
		"requires": {
		  "global-prefix": "^3.0.0"
		},
		"dependencies": {
		  "global-prefix": {
			"version": "3.0.0",
			"resolved": "https://npm.intra.longguikeji.com/global-prefix/-/global-prefix-3.0.0.tgz",
			"integrity": "sha512-awConJSVCHVGND6x3tmMaKcQvwXLhjdkmomy2W+Goaui8YPgYgXJZewhg3fWC+DlfqqQuWg8AwqjGTD2nAPVWg==",
			"dev": true,
			"requires": {
			  "ini": "^1.3.5",
			  "kind-of": "^6.0.2",
			  "which": "^1.3.1"
			}
		  }
		}
	  },
	  "global-prefix": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/global-prefix/-/global-prefix-1.0.2.tgz",
		"integrity": "sha1-2/dDxsFJklk8ZVVoy2btMsASLr4=",
		"dev": true,
		"requires": {
		  "expand-tilde": "^2.0.2",
		  "homedir-polyfill": "^1.0.1",
		  "ini": "^1.3.4",
		  "is-windows": "^1.0.1",
		  "which": "^1.2.14"
		}
	  },
	  "globals": {
		"version": "11.12.0",
		"resolved": "https://npm.intra.longguikeji.com/globals/-/globals-11.12.0.tgz",
		"integrity": "sha512-WOBp/EEGUiIsJSp7wcv/y6MO+lV9UoncWqxuFfm8eBwzWNgyfBd6Gz+IeKQ9jCmyhoH99g15M3T+QaVHFjizVA==",
		"dev": true
	  },
	  "graceful-fs": {
		"version": "4.2.3",
		"resolved": "https://npm.intra.longguikeji.com/graceful-fs/-/graceful-fs-4.2.3.tgz",
		"integrity": "sha512-a30VEBm4PEdx1dRB7MFK7BejejvCvBronbLjht+sHuGYj8PHs7M/5Z+rt5lw551vZ7yfTCj4Vuyy3mSJytDWRQ==",
		"dev": true
	  },
	  "has-flag": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/has-flag/-/has-flag-3.0.0.tgz",
		"integrity": "sha1-tdRU3CGZriJWmfNGfloH87lVuv0=",
		"dev": true
	  },
	  "has-symbols": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/has-symbols/-/has-symbols-1.0.1.tgz",
		"integrity": "sha512-PLcsoqu++dmEIZB+6totNFKq/7Do+Z0u4oT0zKOJNl3lYK6vGwwu2hjHs+68OEZbTjiUE9bgOABXbP/GvrS0Kg==",
		"dev": true
	  },
	  "has-value": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/has-value/-/has-value-1.0.0.tgz",
		"integrity": "sha1-GLKB2lhbHFxR3vJMkw7SmgvmsXc=",
		"dev": true,
		"requires": {
		  "get-value": "^2.0.6",
		  "has-values": "^1.0.0",
		  "isobject": "^3.0.0"
		}
	  },
	  "has-values": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/has-values/-/has-values-1.0.0.tgz",
		"integrity": "sha1-lbC2P+whRmGab+V/51Yo1aOe/k8=",
		"dev": true,
		"requires": {
		  "is-number": "^3.0.0",
		  "kind-of": "^4.0.0"
		},
		"dependencies": {
		  "is-buffer": {
			"version": "1.1.6",
			"resolved": "https://npm.intra.longguikeji.com/is-buffer/-/is-buffer-1.1.6.tgz",
			"integrity": "sha512-NcdALwpXkTm5Zvvbk7owOUSvVvBKDgKP5/ewfXEznmQFfs4ZRmanOeKBTjRVjka3QFoN6XJ+9F3USqfHqTaU5w==",
			"dev": true
		  },
		  "kind-of": {
			"version": "4.0.0",
			"resolved": "https://npm.intra.longguikeji.com/kind-of/-/kind-of-4.0.0.tgz",
			"integrity": "sha1-IIE989cSkosgc3hpGkUGb65y3Vc=",
			"dev": true,
			"requires": {
			  "is-buffer": "^1.1.5"
			}
		  }
		}
	  },
	  "hash-base": {
		"version": "3.0.4",
		"resolved": "https://npm.intra.longguikeji.com/hash-base/-/hash-base-3.0.4.tgz",
		"integrity": "sha1-X8hoaEfs1zSZQDMZprCj8/auSRg=",
		"dev": true,
		"requires": {
		  "inherits": "^2.0.1",
		  "safe-buffer": "^5.0.1"
		}
	  },
	  "hash.js": {
		"version": "1.1.7",
		"resolved": "https://npm.intra.longguikeji.com/hash.js/-/hash.js-1.1.7.tgz",
		"integrity": "sha512-taOaskGt4z4SOANNseOviYDvjEJinIkRgmp7LbKP2YTTmVxWBl87s/uzK9r+44BclBSp2X7K1hqeNfz9JbBeXA==",
		"dev": true,
		"requires": {
		  "inherits": "^2.0.3",
		  "minimalistic-assert": "^1.0.1"
		}
	  },
	  "hmac-drbg": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/hmac-drbg/-/hmac-drbg-1.0.1.tgz",
		"integrity": "sha1-0nRXAQJabHdabFRXk+1QL8DGSaE=",
		"dev": true,
		"requires": {
		  "hash.js": "^1.0.3",
		  "minimalistic-assert": "^1.0.0",
		  "minimalistic-crypto-utils": "^1.0.1"
		}
	  },
	  "homedir-polyfill": {
		"version": "1.0.3",
		"resolved": "https://npm.intra.longguikeji.com/homedir-polyfill/-/homedir-polyfill-1.0.3.tgz",
		"integrity": "sha512-eSmmWE5bZTK2Nou4g0AI3zZ9rswp7GRKoKXS1BLUkvPviOqs4YTN1djQIqrXy9k5gEtdLPy86JjRwsNM9tnDcA==",
		"dev": true,
		"requires": {
		  "parse-passwd": "^1.0.0"
		}
	  },
	  "http-errors": {
		"version": "1.7.2",
		"resolved": "https://npm.intra.longguikeji.com/http-errors/-/http-errors-1.7.2.tgz",
		"integrity": "sha512-uUQBt3H/cSIVfch6i1EuPNy/YsRSOUBXTVfZ+yR7Zjez3qjBz6i9+i4zjNaoqcoFVI4lQJ5plg63TvGfRSDCRg==",
		"requires": {
		  "depd": "~1.1.2",
		  "inherits": "2.0.3",
		  "setprototypeof": "1.1.1",
		  "statuses": ">= 1.5.0 < 2",
		  "toidentifier": "1.0.0"
		}
	  },
	  "https-browserify": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/https-browserify/-/https-browserify-1.0.0.tgz",
		"integrity": "sha1-7AbBDgo0wPL68Zn3/X/Hj//QPHM=",
		"dev": true
	  },
	  "iconv-lite": {
		"version": "0.4.24",
		"resolved": "https://npm.intra.longguikeji.com/iconv-lite/-/iconv-lite-0.4.24.tgz",
		"integrity": "sha512-v3MXnZAcvnywkTUEZomIActle7RXXeedOR31wwl7VlyoXO4Qi9arvSenNQWne1TcRwhCL1HwLI21bEqdpj8/rA==",
		"requires": {
		  "safer-buffer": ">= 2.1.2 < 3"
		}
	  },
	  "ieee754": {
		"version": "1.1.13",
		"resolved": "https://npm.intra.longguikeji.com/ieee754/-/ieee754-1.1.13.tgz",
		"integrity": "sha512-4vf7I2LYV/HaWerSo3XmlMkp5eZ83i+/CDluXi/IGTs/O1sejBNhTtnxzmRZfvOUqj7lZjqHkeTvpgSFDlWZTg==",
		"dev": true
	  },
	  "iferr": {
		"version": "0.1.5",
		"resolved": "https://npm.intra.longguikeji.com/iferr/-/iferr-0.1.5.tgz",
		"integrity": "sha1-xg7taebY/bazEEofy8ocGS3FtQE=",
		"dev": true
	  },
	  "import-local": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/import-local/-/import-local-2.0.0.tgz",
		"integrity": "sha512-b6s04m3O+s3CGSbqDIyP4R6aAwAeYlVq9+WUWep6iHa8ETRf9yei1U48C5MmfJmV9AiLYYBKPMq/W+/WRpQmCQ==",
		"dev": true,
		"requires": {
		  "pkg-dir": "^3.0.0",
		  "resolve-cwd": "^2.0.0"
		}
	  },
	  "imurmurhash": {
		"version": "0.1.4",
		"resolved": "https://npm.intra.longguikeji.com/imurmurhash/-/imurmurhash-0.1.4.tgz",
		"integrity": "sha1-khi5srkoojixPcT7a21XbyMUU+o=",
		"dev": true
	  },
	  "infer-owner": {
		"version": "1.0.4",
		"resolved": "https://npm.intra.longguikeji.com/infer-owner/-/infer-owner-1.0.4.tgz",
		"integrity": "sha512-IClj+Xz94+d7irH5qRyfJonOdfTzuDaifE6ZPWfx0N0+/ATZCbuTPq2prFl526urkQd90WyUKIh1DfBQ2hMz9A==",
		"dev": true
	  },
	  "inflight": {
		"version": "1.0.6",
		"resolved": "https://npm.intra.longguikeji.com/inflight/-/inflight-1.0.6.tgz",
		"integrity": "sha1-Sb1jMdfQLQwJvJEKEHW6gWW1bfk=",
		"dev": true,
		"requires": {
		  "once": "^1.3.0",
		  "wrappy": "1"
		}
	  },
	  "inherits": {
		"version": "2.0.3",
		"resolved": "https://npm.intra.longguikeji.com/inherits/-/inherits-2.0.3.tgz",
		"integrity": "sha1-Yzwsg+PaQqUC9SRmAiSA9CCCYd4="
	  },
	  "ini": {
		"version": "1.3.5",
		"resolved": "https://npm.intra.longguikeji.com/ini/-/ini-1.3.5.tgz",
		"integrity": "sha512-RZY5huIKCMRWDUqZlEi72f/lmXKMvuszcMBduliQ3nnWbx9X/ZBQO7DijMEYS9EhHBb2qacRUMtC7svLwe0lcw==",
		"dev": true
	  },
	  "interpret": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/interpret/-/interpret-1.2.0.tgz",
		"integrity": "sha512-mT34yGKMNceBQUoVn7iCDKDntA7SC6gycMAWzGx1z/CMCTV7b2AAtXlo3nRyHZ1FelRkQbQjprHSYGwzLtkVbw==",
		"dev": true
	  },
	  "invariant": {
		"version": "2.2.4",
		"resolved": "https://npm.intra.longguikeji.com/invariant/-/invariant-2.2.4.tgz",
		"integrity": "sha512-phJfQVBuaJM5raOpJjSfkiD6BpbCE4Ns//LaXl6wGYtUBY83nWS6Rf9tXm2e8VaK60JEjYldbPif/A2B1C2gNA==",
		"dev": true,
		"requires": {
		  "loose-envify": "^1.0.0"
		}
	  },
	  "invert-kv": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/invert-kv/-/invert-kv-2.0.0.tgz",
		"integrity": "sha512-wPVv/y/QQ/Uiirj/vh3oP+1Ww+AWehmi1g5fFWGPF6IpCBCDVrhgHRMvrLfdYcwDh3QJbGXDW4JAuzxElLSqKA==",
		"dev": true
	  },
	  "ipaddr.js": {
		"version": "1.9.0",
		"resolved": "https://npm.intra.longguikeji.com/ipaddr.js/-/ipaddr.js-1.9.0.tgz",
		"integrity": "sha512-M4Sjn6N/+O6/IXSJseKqHoFc+5FdGJ22sXqnjTpdZweHK64MzEPAyQZyEU3R/KRv2GLoa7nNtg/C2Ev6m7z+eA=="
	  },
	  "is-accessor-descriptor": {
		"version": "0.1.6",
		"resolved": "https://npm.intra.longguikeji.com/is-accessor-descriptor/-/is-accessor-descriptor-0.1.6.tgz",
		"integrity": "sha1-qeEss66Nh2cn7u84Q/igiXtcmNY=",
		"dev": true,
		"requires": {
		  "kind-of": "^3.0.2"
		},
		"dependencies": {
		  "is-buffer": {
			"version": "1.1.6",
			"resolved": "https://npm.intra.longguikeji.com/is-buffer/-/is-buffer-1.1.6.tgz",
			"integrity": "sha512-NcdALwpXkTm5Zvvbk7owOUSvVvBKDgKP5/ewfXEznmQFfs4ZRmanOeKBTjRVjka3QFoN6XJ+9F3USqfHqTaU5w==",
			"dev": true
		  },
		  "kind-of": {
			"version": "3.2.2",
			"resolved": "https://npm.intra.longguikeji.com/kind-of/-/kind-of-3.2.2.tgz",
			"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
			"dev": true,
			"requires": {
			  "is-buffer": "^1.1.5"
			}
		  }
		}
	  },
	  "is-binary-path": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/is-binary-path/-/is-binary-path-1.0.1.tgz",
		"integrity": "sha1-dfFmQrSA8YenEcgUFh/TpKdlWJg=",
		"dev": true,
		"requires": {
		  "binary-extensions": "^1.0.0"
		}
	  },
	  "is-buffer": {
		"version": "2.0.4",
		"resolved": "https://npm.intra.longguikeji.com/is-buffer/-/is-buffer-2.0.4.tgz",
		"integrity": "sha512-Kq1rokWXOPXWuaMAqZiJW4XxsmD9zGx9q4aePabbn3qCRGedtH7Cm+zV8WETitMfu1wdh+Rvd6w5egwSngUX2A=="
	  },
	  "is-data-descriptor": {
		"version": "0.1.4",
		"resolved": "https://npm.intra.longguikeji.com/is-data-descriptor/-/is-data-descriptor-0.1.4.tgz",
		"integrity": "sha1-C17mSDiOLIYCgueT8YVv7D8wG1Y=",
		"dev": true,
		"requires": {
		  "kind-of": "^3.0.2"
		},
		"dependencies": {
		  "is-buffer": {
			"version": "1.1.6",
			"resolved": "https://npm.intra.longguikeji.com/is-buffer/-/is-buffer-1.1.6.tgz",
			"integrity": "sha512-NcdALwpXkTm5Zvvbk7owOUSvVvBKDgKP5/ewfXEznmQFfs4ZRmanOeKBTjRVjka3QFoN6XJ+9F3USqfHqTaU5w==",
			"dev": true
		  },
		  "kind-of": {
			"version": "3.2.2",
			"resolved": "https://npm.intra.longguikeji.com/kind-of/-/kind-of-3.2.2.tgz",
			"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
			"dev": true,
			"requires": {
			  "is-buffer": "^1.1.5"
			}
		  }
		}
	  },
	  "is-descriptor": {
		"version": "0.1.6",
		"resolved": "https://npm.intra.longguikeji.com/is-descriptor/-/is-descriptor-0.1.6.tgz",
		"integrity": "sha512-avDYr0SB3DwO9zsMov0gKCESFYqCnE4hq/4z3TdUlukEy5t9C0YRq7HLrsN52NAcqXKaepeCD0n+B0arnVG3Hg==",
		"dev": true,
		"requires": {
		  "is-accessor-descriptor": "^0.1.6",
		  "is-data-descriptor": "^0.1.4",
		  "kind-of": "^5.0.0"
		},
		"dependencies": {
		  "kind-of": {
			"version": "5.1.0",
			"resolved": "https://npm.intra.longguikeji.com/kind-of/-/kind-of-5.1.0.tgz",
			"integrity": "sha512-NGEErnH6F2vUuXDh+OlbcKW7/wOcfdRHaZ7VWtqCztfHri/++YKmP51OdWeGPuqCOba6kk2OTe5d02VmTB80Pw==",
			"dev": true
		  }
		}
	  },
	  "is-extendable": {
		"version": "0.1.1",
		"resolved": "https://npm.intra.longguikeji.com/is-extendable/-/is-extendable-0.1.1.tgz",
		"integrity": "sha1-YrEQ4omkcUGOPsNqYX1HLjAd/Ik=",
		"dev": true
	  },
	  "is-extglob": {
		"version": "2.1.1",
		"resolved": "https://npm.intra.longguikeji.com/is-extglob/-/is-extglob-2.1.1.tgz",
		"integrity": "sha1-qIwCU1eR8C7TfHahueqXc8gz+MI=",
		"dev": true
	  },
	  "is-fullwidth-code-point": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/is-fullwidth-code-point/-/is-fullwidth-code-point-3.0.0.tgz",
		"integrity": "sha512-zymm5+u+sCsSWyD9qNaejV3DFvhCKclKdizYaJUuHA83RLjb7nSuGnddCHGv0hk+KY7BMAlsWeK4Ueg6EV6XQg=="
	  },
	  "is-glob": {
		"version": "4.0.1",
		"resolved": "https://npm.intra.longguikeji.com/is-glob/-/is-glob-4.0.1.tgz",
		"integrity": "sha512-5G0tKtBTFImOqDnLB2hG6Bp2qcKEFduo4tZu9MT/H6NQv/ghhy30o55ufafxJ/LdH79LLs2Kfrn85TLKyA7BUg==",
		"dev": true,
		"requires": {
		  "is-extglob": "^2.1.1"
		}
	  },
	  "is-number": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/is-number/-/is-number-3.0.0.tgz",
		"integrity": "sha1-JP1iAaR4LPUFYcgQJ2r8fRLXEZU=",
		"dev": true,
		"requires": {
		  "kind-of": "^3.0.2"
		},
		"dependencies": {
		  "is-buffer": {
			"version": "1.1.6",
			"resolved": "https://npm.intra.longguikeji.com/is-buffer/-/is-buffer-1.1.6.tgz",
			"integrity": "sha512-NcdALwpXkTm5Zvvbk7owOUSvVvBKDgKP5/ewfXEznmQFfs4ZRmanOeKBTjRVjka3QFoN6XJ+9F3USqfHqTaU5w==",
			"dev": true
		  },
		  "kind-of": {
			"version": "3.2.2",
			"resolved": "https://npm.intra.longguikeji.com/kind-of/-/kind-of-3.2.2.tgz",
			"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
			"dev": true,
			"requires": {
			  "is-buffer": "^1.1.5"
			}
		  }
		}
	  },
	  "is-plain-object": {
		"version": "2.0.4",
		"resolved": "https://npm.intra.longguikeji.com/is-plain-object/-/is-plain-object-2.0.4.tgz",
		"integrity": "sha512-h5PpgXkWitc38BBMYawTYMWJHFZJVnBquFE57xFpjB8pJFiF6gZ+bU+WyI/yqXiFR5mdLsgYNaPe8uao6Uv9Og==",
		"dev": true,
		"requires": {
		  "isobject": "^3.0.1"
		}
	  },
	  "is-stream": {
		"version": "1.1.0",
		"resolved": "https://npm.intra.longguikeji.com/is-stream/-/is-stream-1.1.0.tgz",
		"integrity": "sha1-EtSj3U5o4Lec6428hBc66A2RykQ=",
		"dev": true
	  },
	  "is-windows": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/is-windows/-/is-windows-1.0.2.tgz",
		"integrity": "sha512-eXK1UInq2bPmjyX6e3VHIzMLobc4J94i4AWn+Hpq3OU5KkrRC96OAcR3PRJ/pGu6m8TRnBHP9dkXQVsT/COVIA==",
		"dev": true
	  },
	  "is-wsl": {
		"version": "1.1.0",
		"resolved": "https://npm.intra.longguikeji.com/is-wsl/-/is-wsl-1.1.0.tgz",
		"integrity": "sha1-HxbkqiKwTRM2tmGIpmrzxgDDpm0=",
		"dev": true
	  },
	  "isarray": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/isarray/-/isarray-1.0.0.tgz",
		"integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
		"dev": true
	  },
	  "isexe": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/isexe/-/isexe-2.0.0.tgz",
		"integrity": "sha1-6PvzdNxVb/iUehDcsFctYz8s+hA=",
		"dev": true
	  },
	  "isobject": {
		"version": "3.0.1",
		"resolved": "https://npm.intra.longguikeji.com/isobject/-/isobject-3.0.1.tgz",
		"integrity": "sha1-TkMekrEalzFjaqH5yNHMvP2reN8=",
		"dev": true
	  },
	  "js-levenshtein": {
		"version": "1.1.6",
		"resolved": "https://npm.intra.longguikeji.com/js-levenshtein/-/js-levenshtein-1.1.6.tgz",
		"integrity": "sha512-X2BB11YZtrRqY4EnQcLX5Rh373zbK4alC1FW7D7MBhL2gtcC17cTnr6DmfHZeS0s2rTHjUTMMHfG7gO8SSdw+g==",
		"dev": true
	  },
	  "js-tokens": {
		"version": "4.0.0",
		"resolved": "https://npm.intra.longguikeji.com/js-tokens/-/js-tokens-4.0.0.tgz",
		"integrity": "sha512-RdJUflcE3cUzKiMqQgsCu06FPu9UdIJO0beYbPhHN4k6apgJtifcoCtT9bcxOpYBtpD2kCM6Sbzg4CausW/PKQ==",
		"dev": true
	  },
	  "jsesc": {
		"version": "2.5.2",
		"resolved": "https://npm.intra.longguikeji.com/jsesc/-/jsesc-2.5.2.tgz",
		"integrity": "sha512-OYu7XEzjkCQ3C5Ps3QIZsQfNpqoJyZZA99wd9aWd05NCtC5pWOkShK2mkL6HXQR6/Cy2lbNdPlZBpuQHXE63gA==",
		"dev": true
	  },
	  "json-parse-better-errors": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/json-parse-better-errors/-/json-parse-better-errors-1.0.2.tgz",
		"integrity": "sha512-mrqyZKfX5EhL7hvqcV6WG1yYjnjeuYDzDhhcAAUrq8Po85NBQBJP+ZDUT75qZQ98IkUoBqdkExkukOU7Ts2wrw==",
		"dev": true
	  },
	  "json-schema-traverse": {
		"version": "0.4.1",
		"resolved": "https://npm.intra.longguikeji.com/json-schema-traverse/-/json-schema-traverse-0.4.1.tgz",
		"integrity": "sha512-xbbCH5dCYU5T8LcEhhuh7HJ88HXuW3qsI3Y0zOZFKfZEHcpWiHU/Jxzk629Brsab/mMiHQti9wMP+845RPe3Vg==",
		"dev": true
	  },
	  "json5": {
		"version": "2.1.1",
		"resolved": "https://npm.intra.longguikeji.com/json5/-/json5-2.1.1.tgz",
		"integrity": "sha512-l+3HXD0GEI3huGq1njuqtzYK8OYJyXMkOLtQ53pjWh89tvWS2h6l+1zMkYWqlb57+SiQodKZyvMEFb2X+KrFhQ==",
		"dev": true,
		"requires": {
		  "minimist": "^1.2.0"
		}
	  },
	  "kind-of": {
		"version": "6.0.2",
		"resolved": "https://npm.intra.longguikeji.com/kind-of/-/kind-of-6.0.2.tgz",
		"integrity": "sha512-s5kLOcnH0XqDO+FvuaLX8DDjZ18CGFk7VygH40QoKPUQhW4e2rvM0rwUq0t8IQDOwYSeLK01U90OjzBTme2QqA==",
		"dev": true
	  },
	  "lcid": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/lcid/-/lcid-2.0.0.tgz",
		"integrity": "sha512-avPEb8P8EGnwXKClwsNUgryVjllcRqtMYa49NTsbQagYuT1DcXnl1915oxWjoyGrXR6zH/Y0Zc96xWsPcoDKeA==",
		"dev": true,
		"requires": {
		  "invert-kv": "^2.0.0"
		}
	  },
	  "loadash": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/loadash/-/loadash-1.0.0.tgz",
		"integrity": "sha512-xlX5HBsXB3KG0FJbJJG/3kYWCfsCyCSus3T+uHVu6QL6YxAdggmm3QeyLgn54N2yi5/UE6xxL5ZWJAAiHzHYEg=="
	  },
	  "loader-runner": {
		"version": "2.4.0",
		"resolved": "https://npm.intra.longguikeji.com/loader-runner/-/loader-runner-2.4.0.tgz",
		"integrity": "sha512-Jsmr89RcXGIwivFY21FcRrisYZfvLMTWx5kOLc+JTxtpBOG6xML0vzbc6SEQG2FO9/4Fc3wW4LVcB5DmGflaRw==",
		"dev": true
	  },
	  "loader-utils": {
		"version": "1.2.3",
		"resolved": "https://npm.intra.longguikeji.com/loader-utils/-/loader-utils-1.2.3.tgz",
		"integrity": "sha512-fkpz8ejdnEMG3s37wGL07iSBDg99O9D5yflE9RGNH3hRdx9SOwYfnGYdZOUIZitN8E+E2vkq3MUMYMvPYl5ZZA==",
		"dev": true,
		"requires": {
		  "big.js": "^5.2.2",
		  "emojis-list": "^2.0.0",
		  "json5": "^1.0.1"
		},
		"dependencies": {
		  "json5": {
			"version": "1.0.1",
			"resolved": "https://npm.intra.longguikeji.com/json5/-/json5-1.0.1.tgz",
			"integrity": "sha512-aKS4WQjPenRxiQsC93MNfjx+nbF4PAdYzmd/1JIj8HYzqfbu86beTuNgXDzPknWk0n0uARlyewZo4s++ES36Ow==",
			"dev": true,
			"requires": {
			  "minimist": "^1.2.0"
			}
		  }
		}
	  },
	  "locate-path": {
		"version": "5.0.0",
		"resolved": "https://npm.intra.longguikeji.com/locate-path/-/locate-path-5.0.0.tgz",
		"integrity": "sha512-t7hw9pI+WvuwNJXwk5zVHpyhIqzg2qTlklJOf0mVxGSbe3Fp2VieZcduNYjaLDoy6p9uGpQEGWG87WpMKlNq8g==",
		"requires": {
		  "p-locate": "^4.1.0"
		}
	  },
	  "lodash": {
		"version": "4.17.15",
		"resolved": "https://npm.intra.longguikeji.com/lodash/-/lodash-4.17.15.tgz",
		"integrity": "sha512-8xOcRHvCjnocdS5cpwXQXVzmmh5e5+saE2QGoeQmbKmRS6J3VQppPOIt0MnmE+4xlZoumy0GPG0D0MVIQbNA1A==",
		"dev": true
	  },
	  "loose-envify": {
		"version": "1.4.0",
		"resolved": "https://npm.intra.longguikeji.com/loose-envify/-/loose-envify-1.4.0.tgz",
		"integrity": "sha512-lyuxPGr/Wfhrlem2CL/UcnUc1zcqKAImBDzukY7Y5F/yQiNdko6+fRLevlw1HgMySw7f611UIY408EtxRSoK3Q==",
		"dev": true,
		"requires": {
		  "js-tokens": "^3.0.0 || ^4.0.0"
		}
	  },
	  "lru-cache": {
		"version": "5.1.1",
		"resolved": "https://npm.intra.longguikeji.com/lru-cache/-/lru-cache-5.1.1.tgz",
		"integrity": "sha512-KpNARQA3Iwv+jTA0utUVVbrh+Jlrr1Fv0e56GGzAFOXN7dk/FviaDW8LHmK52DlcH4WP2n6gI8vN1aesBFgo9w==",
		"dev": true,
		"requires": {
		  "yallist": "^3.0.2"
		}
	  },
	  "make-dir": {
		"version": "2.1.0",
		"resolved": "https://npm.intra.longguikeji.com/make-dir/-/make-dir-2.1.0.tgz",
		"integrity": "sha512-LS9X+dc8KLxXCb8dni79fLIIUA5VyZoyjSMCwTluaXA0o27cCK0bhXkpgw+sTXVpPy/lSO57ilRixqk0vDmtRA==",
		"dev": true,
		"requires": {
		  "pify": "^4.0.1",
		  "semver": "^5.6.0"
		}
	  },
	  "mamacro": {
		"version": "0.0.3",
		"resolved": "https://npm.intra.longguikeji.com/mamacro/-/mamacro-0.0.3.tgz",
		"integrity": "sha512-qMEwh+UujcQ+kbz3T6V+wAmO2U8veoq2w+3wY8MquqwVA3jChfwY+Tk52GZKDfACEPjuZ7r2oJLejwpt8jtwTA==",
		"dev": true
	  },
	  "map-age-cleaner": {
		"version": "0.1.3",
		"resolved": "https://npm.intra.longguikeji.com/map-age-cleaner/-/map-age-cleaner-0.1.3.tgz",
		"integrity": "sha512-bJzx6nMoP6PDLPBFmg7+xRKeFZvFboMrGlxmNj9ClvX53KrmvM5bXFXEWjbz4cz1AFn+jWJ9z/DJSz7hrs0w3w==",
		"dev": true,
		"requires": {
		  "p-defer": "^1.0.0"
		}
	  },
	  "map-cache": {
		"version": "0.2.2",
		"resolved": "https://npm.intra.longguikeji.com/map-cache/-/map-cache-0.2.2.tgz",
		"integrity": "sha1-wyq9C9ZSXZsFFkW7TyasXcmKDb8=",
		"dev": true
	  },
	  "map-visit": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/map-visit/-/map-visit-1.0.0.tgz",
		"integrity": "sha1-7Nyo8TFE5mDxtb1B8S80edmN+48=",
		"dev": true,
		"requires": {
		  "object-visit": "^1.0.0"
		}
	  },
	  "md5.js": {
		"version": "1.3.5",
		"resolved": "https://npm.intra.longguikeji.com/md5.js/-/md5.js-1.3.5.tgz",
		"integrity": "sha512-xitP+WxNPcTTOgnTJcrhM0xvdPepipPSf3I8EIpGKeFLjt3PlJLIDG3u8EX53ZIubkb+5U2+3rELYpEhHhzdkg==",
		"dev": true,
		"requires": {
		  "hash-base": "^3.0.0",
		  "inherits": "^2.0.1",
		  "safe-buffer": "^5.1.2"
		}
	  },
	  "media-typer": {
		"version": "0.3.0",
		"resolved": "https://npm.intra.longguikeji.com/media-typer/-/media-typer-0.3.0.tgz",
		"integrity": "sha1-hxDXrwqmJvj/+hzgAWhUUmMlV0g="
	  },
	  "mem": {
		"version": "4.3.0",
		"resolved": "https://npm.intra.longguikeji.com/mem/-/mem-4.3.0.tgz",
		"integrity": "sha512-qX2bG48pTqYRVmDB37rn/6PT7LcR8T7oAX3bf99u1Tt1nzxYfxkgqDwUwolPlXweM0XzBOBFzSx4kfp7KP1s/w==",
		"dev": true,
		"requires": {
		  "map-age-cleaner": "^0.1.1",
		  "mimic-fn": "^2.0.0",
		  "p-is-promise": "^2.0.0"
		}
	  },
	  "memory-fs": {
		"version": "0.4.1",
		"resolved": "https://npm.intra.longguikeji.com/memory-fs/-/memory-fs-0.4.1.tgz",
		"integrity": "sha1-OpoguEYlI+RHz7x+i7gO1me/xVI=",
		"dev": true,
		"requires": {
		  "errno": "^0.1.3",
		  "readable-stream": "^2.0.1"
		}
	  },
	  "merge-descriptors": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/merge-descriptors/-/merge-descriptors-1.0.1.tgz",
		"integrity": "sha1-sAqqVW3YtEVoFQ7J0blT8/kMu2E="
	  },
	  "methods": {
		"version": "1.1.2",
		"resolved": "https://npm.intra.longguikeji.com/methods/-/methods-1.1.2.tgz",
		"integrity": "sha1-VSmk1nZUE07cxSZmVoNbD4Ua/O4="
	  },
	  "micromatch": {
		"version": "3.1.10",
		"resolved": "https://npm.intra.longguikeji.com/micromatch/-/micromatch-3.1.10.tgz",
		"integrity": "sha512-MWikgl9n9M3w+bpsY3He8L+w9eF9338xRl8IAO5viDizwSzziFEyUzo2xrrloB64ADbTf8uA8vRqqttDTOmccg==",
		"dev": true,
		"requires": {
		  "arr-diff": "^4.0.0",
		  "array-unique": "^0.3.2",
		  "braces": "^2.3.1",
		  "define-property": "^2.0.2",
		  "extend-shallow": "^3.0.2",
		  "extglob": "^2.0.4",
		  "fragment-cache": "^0.2.1",
		  "kind-of": "^6.0.2",
		  "nanomatch": "^1.2.9",
		  "object.pick": "^1.3.0",
		  "regex-not": "^1.0.0",
		  "snapdragon": "^0.8.1",
		  "to-regex": "^3.0.2"
		}
	  },
	  "miller-rabin": {
		"version": "4.0.1",
		"resolved": "https://npm.intra.longguikeji.com/miller-rabin/-/miller-rabin-4.0.1.tgz",
		"integrity": "sha512-115fLhvZVqWwHPbClyntxEVfVDfl9DLLTuJvq3g2O/Oxi8AiNouAHvDSzHS0viUJc+V5vm3eq91Xwqn9dp4jRA==",
		"dev": true,
		"requires": {
		  "bn.js": "^4.0.0",
		  "brorand": "^1.0.1"
		}
	  },
	  "mime": {
		"version": "1.6.0",
		"resolved": "https://npm.intra.longguikeji.com/mime/-/mime-1.6.0.tgz",
		"integrity": "sha512-x0Vn8spI+wuJ1O6S7gnbaQg8Pxh4NNHb7KSINmEWKiPE4RKOplvijn+NkmYmmRgP68mc70j2EbeTFRsrswaQeg=="
	  },
	  "mime-db": {
		"version": "1.43.0",
		"resolved": "https://npm.intra.longguikeji.com/mime-db/-/mime-db-1.43.0.tgz",
		"integrity": "sha512-+5dsGEEovYbT8UY9yD7eE4XTc4UwJ1jBYlgaQQF38ENsKR3wj/8q8RFZrF9WIZpB2V1ArTVFUva8sAul1NzRzQ=="
	  },
	  "mime-types": {
		"version": "2.1.26",
		"resolved": "https://npm.intra.longguikeji.com/mime-types/-/mime-types-2.1.26.tgz",
		"integrity": "sha512-01paPWYgLrkqAyrlDorC1uDwl2p3qZT7yl806vW7DvDoxwXi46jsjFbg+WdwotBIk6/MbEhO/dh5aZ5sNj/dWQ==",
		"requires": {
		  "mime-db": "1.43.0"
		}
	  },
	  "mimic-fn": {
		"version": "2.1.0",
		"resolved": "https://npm.intra.longguikeji.com/mimic-fn/-/mimic-fn-2.1.0.tgz",
		"integrity": "sha512-OqbOk5oEQeAZ8WXWydlu9HJjz9WVdEIvamMCcXmuqUYjTknH/sqsWvhQ3vgwKFRR1HpjvNBKQ37nbJgYzGqGcg==",
		"dev": true
	  },
	  "minimalistic-assert": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/minimalistic-assert/-/minimalistic-assert-1.0.1.tgz",
		"integrity": "sha512-UtJcAD4yEaGtjPezWuO9wC4nwUnVH/8/Im3yEHQP4b67cXlD/Qr9hdITCU1xDbSEXg2XKNaP8jsReV7vQd00/A==",
		"dev": true
	  },
	  "minimalistic-crypto-utils": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/minimalistic-crypto-utils/-/minimalistic-crypto-utils-1.0.1.tgz",
		"integrity": "sha1-9sAMHAsIIkblxNmd+4x8CDsrWCo=",
		"dev": true
	  },
	  "minimatch": {
		"version": "3.0.4",
		"resolved": "https://npm.intra.longguikeji.com/minimatch/-/minimatch-3.0.4.tgz",
		"integrity": "sha512-yJHVQEhyqPLUTgt9B83PXu6W3rx4MvvHvSUvToogpwoGDOUQ+yDrR0HRot+yOCdCO7u4hX3pWft6kWBBcqh0UA==",
		"dev": true,
		"requires": {
		  "brace-expansion": "^1.1.7"
		}
	  },
	  "minimist": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/minimist/-/minimist-1.2.0.tgz",
		"integrity": "sha1-o1AIsg9BOD7sH7kU9M1d95omQoQ=",
		"dev": true
	  },
	  "mississippi": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/mississippi/-/mississippi-3.0.0.tgz",
		"integrity": "sha512-x471SsVjUtBRtcvd4BzKE9kFC+/2TeWgKCgw0bZcw1b9l2X3QX5vCWgF+KaZaYm87Ss//rHnWryupDrgLvmSkA==",
		"dev": true,
		"requires": {
		  "concat-stream": "^1.5.0",
		  "duplexify": "^3.4.2",
		  "end-of-stream": "^1.1.0",
		  "flush-write-stream": "^1.0.0",
		  "from2": "^2.1.0",
		  "parallel-transform": "^1.1.0",
		  "pump": "^3.0.0",
		  "pumpify": "^1.3.3",
		  "stream-each": "^1.1.0",
		  "through2": "^2.0.0"
		}
	  },
	  "mixin-deep": {
		"version": "1.3.2",
		"resolved": "https://npm.intra.longguikeji.com/mixin-deep/-/mixin-deep-1.3.2.tgz",
		"integrity": "sha512-WRoDn//mXBiJ1H40rqa3vH0toePwSsGb45iInWlTySa+Uu4k3tYUSxa2v1KqAiLtvlrSzaExqS1gtk96A9zvEA==",
		"dev": true,
		"requires": {
		  "for-in": "^1.0.2",
		  "is-extendable": "^1.0.1"
		},
		"dependencies": {
		  "is-extendable": {
			"version": "1.0.1",
			"resolved": "https://npm.intra.longguikeji.com/is-extendable/-/is-extendable-1.0.1.tgz",
			"integrity": "sha512-arnXMxT1hhoKo9k1LZdmlNyJdDDfy2v0fXjFlmok4+i8ul/6WlbVge9bhM74OpNPQPMGUToDtz+KXa1PneJxOA==",
			"dev": true,
			"requires": {
			  "is-plain-object": "^2.0.4"
			}
		  }
		}
	  },
	  "mkdirp": {
		"version": "0.5.1",
		"resolved": "https://npm.intra.longguikeji.com/mkdirp/-/mkdirp-0.5.1.tgz",
		"integrity": "sha1-MAV0OOrGz3+MR2fzhkjWaX11yQM=",
		"dev": true,
		"requires": {
		  "minimist": "0.0.8"
		},
		"dependencies": {
		  "minimist": {
			"version": "0.0.8",
			"resolved": "https://npm.intra.longguikeji.com/minimist/-/minimist-0.0.8.tgz",
			"integrity": "sha1-hX/Kv8M5fSYluCKCYuhqp6ARsF0=",
			"dev": true
		  }
		}
	  },
	  "module-alias": {
		"version": "2.2.2",
		"resolved": "https://npm.intra.longguikeji.com/module-alias/-/module-alias-2.2.2.tgz",
		"integrity": "sha512-A/78XjoX2EmNvppVWEhM2oGk3x4lLxnkEA4jTbaK97QKSDjkIoOsKQlfylt/d3kKKi596Qy3NP5XrXJ6fZIC9Q=="
	  },
	  "move-concurrently": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/move-concurrently/-/move-concurrently-1.0.1.tgz",
		"integrity": "sha1-viwAX9oy4LKa8fBdfEszIUxwH5I=",
		"dev": true,
		"requires": {
		  "aproba": "^1.1.1",
		  "copy-concurrently": "^1.0.0",
		  "fs-write-stream-atomic": "^1.0.8",
		  "mkdirp": "^0.5.1",
		  "rimraf": "^2.5.4",
		  "run-queue": "^1.0.3"
		}
	  },
	  "ms": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/ms/-/ms-2.0.0.tgz",
		"integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g="
	  },
	  "nan": {
		"version": "2.14.0",
		"resolved": "https://npm.intra.longguikeji.com/nan/-/nan-2.14.0.tgz",
		"integrity": "sha512-INOFj37C7k3AfaNTtX8RhsTw7qRy7eLET14cROi9+5HAVbbHuIWUHEauBv5qT4Av2tWasiTY1Jw6puUNqRJXQg==",
		"dev": true,
		"optional": true
	  },
	  "nanomatch": {
		"version": "1.2.13",
		"resolved": "https://npm.intra.longguikeji.com/nanomatch/-/nanomatch-1.2.13.tgz",
		"integrity": "sha512-fpoe2T0RbHwBTBUOftAfBPaDEi06ufaUai0mE6Yn1kacc3SnTErfb/h+X94VXzI64rKFHYImXSvdwGGCmwOqCA==",
		"dev": true,
		"requires": {
		  "arr-diff": "^4.0.0",
		  "array-unique": "^0.3.2",
		  "define-property": "^2.0.2",
		  "extend-shallow": "^3.0.2",
		  "fragment-cache": "^0.2.1",
		  "is-windows": "^1.0.2",
		  "kind-of": "^6.0.2",
		  "object.pick": "^1.3.0",
		  "regex-not": "^1.0.0",
		  "snapdragon": "^0.8.1",
		  "to-regex": "^3.0.1"
		}
	  },
	  "negotiator": {
		"version": "0.6.2",
		"resolved": "https://npm.intra.longguikeji.com/negotiator/-/negotiator-0.6.2.tgz",
		"integrity": "sha512-hZXc7K2e+PgeI1eDBe/10Ard4ekbfrrqG8Ep+8Jmf4JID2bNg7NvCPOZN+kfF574pFQI7mum2AUqDidoKqcTOw=="
	  },
	  "neo-async": {
		"version": "2.6.1",
		"resolved": "https://npm.intra.longguikeji.com/neo-async/-/neo-async-2.6.1.tgz",
		"integrity": "sha512-iyam8fBuCUpWeKPGpaNMetEocMt364qkCsfL9JuhjXX6dRnguRVOfk2GZaDpPjcOKiiXCPINZC1GczQ7iTq3Zw==",
		"dev": true
	  },
	  "nice-try": {
		"version": "1.0.5",
		"resolved": "https://npm.intra.longguikeji.com/nice-try/-/nice-try-1.0.5.tgz",
		"integrity": "sha512-1nh45deeb5olNY7eX82BkPO7SSxR5SSYJiPTrTdFUVYwAl8CKMA5N9PjTYkHiRjisVcxcQ1HXdLhx2qxxJzLNQ==",
		"dev": true
	  },
	  "node-libs-browser": {
		"version": "2.2.1",
		"resolved": "https://npm.intra.longguikeji.com/node-libs-browser/-/node-libs-browser-2.2.1.tgz",
		"integrity": "sha512-h/zcD8H9kaDZ9ALUWwlBUDo6TKF8a7qBSCSEGfjTVIYeqsioSKaAX+BN7NgiMGp6iSIXZ3PxgCu8KS3b71YK5Q==",
		"dev": true,
		"requires": {
		  "assert": "^1.1.1",
		  "browserify-zlib": "^0.2.0",
		  "buffer": "^4.3.0",
		  "console-browserify": "^1.1.0",
		  "constants-browserify": "^1.0.0",
		  "crypto-browserify": "^3.11.0",
		  "domain-browser": "^1.1.1",
		  "events": "^3.0.0",
		  "https-browserify": "^1.0.0",
		  "os-browserify": "^0.3.0",
		  "path-browserify": "0.0.1",
		  "process": "^0.11.10",
		  "punycode": "^1.2.4",
		  "querystring-es3": "^0.2.0",
		  "readable-stream": "^2.3.3",
		  "stream-browserify": "^2.0.1",
		  "stream-http": "^2.7.2",
		  "string_decoder": "^1.0.0",
		  "timers-browserify": "^2.0.4",
		  "tty-browserify": "0.0.0",
		  "url": "^0.11.0",
		  "util": "^0.11.0",
		  "vm-browserify": "^1.0.1"
		},
		"dependencies": {
		  "punycode": {
			"version": "1.4.1",
			"resolved": "https://npm.intra.longguikeji.com/punycode/-/punycode-1.4.1.tgz",
			"integrity": "sha1-wNWmOycYgArY4esPpSachN1BhF4=",
			"dev": true
		  }
		}
	  },
	  "node-releases": {
		"version": "1.1.44",
		"resolved": "https://npm.intra.longguikeji.com/node-releases/-/node-releases-1.1.44.tgz",
		"integrity": "sha512-NwbdvJyR7nrcGrXvKAvzc5raj/NkoJudkarh2yIpJ4t0NH4aqjUDz/486P+ynIW5eokKOfzGNRdYoLfBlomruw==",
		"dev": true,
		"requires": {
		  "semver": "^6.3.0"
		},
		"dependencies": {
		  "semver": {
			"version": "6.3.0",
			"resolved": "https://npm.intra.longguikeji.com/semver/-/semver-6.3.0.tgz",
			"integrity": "sha512-b39TBaTSfV6yBrapU89p5fKekE2m/NwnDocOVruQFS1/veMgdzuPcnOM34M6CwxW8jH/lxEa5rBoDeUwu5HHTw==",
			"dev": true
		  }
		}
	  },
	  "normalize-path": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/normalize-path/-/normalize-path-3.0.0.tgz",
		"integrity": "sha512-6eZs5Ls3WtCisHWp9S2GUy8dqkpGi4BVSz3GaqiE6ezub0512ESztXUwUB6C6IKbQkY2Pnb/mD4WYojCRwcwLA==",
		"dev": true
	  },
	  "npm-run-path": {
		"version": "2.0.2",
		"resolved": "https://npm.intra.longguikeji.com/npm-run-path/-/npm-run-path-2.0.2.tgz",
		"integrity": "sha1-NakjLfo11wZ7TLLd8jV7GHFTbF8=",
		"dev": true,
		"requires": {
		  "path-key": "^2.0.0"
		}
	  },
	  "object-assign": {
		"version": "4.1.1",
		"resolved": "https://npm.intra.longguikeji.com/object-assign/-/object-assign-4.1.1.tgz",
		"integrity": "sha1-IQmtx5ZYh8/AXLvUQsrIv7s2CGM=",
		"dev": true
	  },
	  "object-copy": {
		"version": "0.1.0",
		"resolved": "https://npm.intra.longguikeji.com/object-copy/-/object-copy-0.1.0.tgz",
		"integrity": "sha1-fn2Fi3gb18mRpBupde04EnVOmYw=",
		"dev": true,
		"requires": {
		  "copy-descriptor": "^0.1.0",
		  "define-property": "^0.2.5",
		  "kind-of": "^3.0.3"
		},
		"dependencies": {
		  "define-property": {
			"version": "0.2.5",
			"resolved": "https://npm.intra.longguikeji.com/define-property/-/define-property-0.2.5.tgz",
			"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
			"dev": true,
			"requires": {
			  "is-descriptor": "^0.1.0"
			}
		  },
		  "is-buffer": {
			"version": "1.1.6",
			"resolved": "https://npm.intra.longguikeji.com/is-buffer/-/is-buffer-1.1.6.tgz",
			"integrity": "sha512-NcdALwpXkTm5Zvvbk7owOUSvVvBKDgKP5/ewfXEznmQFfs4ZRmanOeKBTjRVjka3QFoN6XJ+9F3USqfHqTaU5w==",
			"dev": true
		  },
		  "kind-of": {
			"version": "3.2.2",
			"resolved": "https://npm.intra.longguikeji.com/kind-of/-/kind-of-3.2.2.tgz",
			"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
			"dev": true,
			"requires": {
			  "is-buffer": "^1.1.5"
			}
		  }
		}
	  },
	  "object-keys": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/object-keys/-/object-keys-1.1.1.tgz",
		"integrity": "sha512-NuAESUOUMrlIXOfHKzD6bpPu3tYt3xvjNdRIQ+FeT0lNb4K8WR70CaDxhuNguS2XG+GjkyMwOzsN5ZktImfhLA==",
		"dev": true
	  },
	  "object-visit": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/object-visit/-/object-visit-1.0.1.tgz",
		"integrity": "sha1-95xEk68MU3e1n+OdOV5BBC3QRbs=",
		"dev": true,
		"requires": {
		  "isobject": "^3.0.0"
		}
	  },
	  "object.assign": {
		"version": "4.1.0",
		"resolved": "https://npm.intra.longguikeji.com/object.assign/-/object.assign-4.1.0.tgz",
		"integrity": "sha512-exHJeq6kBKj58mqGyTQ9DFvrZC/eR6OwxzoM9YRoGBqrXYonaFyGiFMuc9VZrXf7DarreEwMpurG3dd+CNyW5w==",
		"dev": true,
		"requires": {
		  "define-properties": "^1.1.2",
		  "function-bind": "^1.1.1",
		  "has-symbols": "^1.0.0",
		  "object-keys": "^1.0.11"
		}
	  },
	  "object.pick": {
		"version": "1.3.0",
		"resolved": "https://npm.intra.longguikeji.com/object.pick/-/object.pick-1.3.0.tgz",
		"integrity": "sha1-h6EKxMFpS9Lhy/U1kaZhQftd10c=",
		"dev": true,
		"requires": {
		  "isobject": "^3.0.1"
		}
	  },
	  "on-finished": {
		"version": "2.3.0",
		"resolved": "https://npm.intra.longguikeji.com/on-finished/-/on-finished-2.3.0.tgz",
		"integrity": "sha1-IPEzZIGwg811M3mSoWlxqi2QaUc=",
		"requires": {
		  "ee-first": "1.1.1"
		}
	  },
	  "once": {
		"version": "1.4.0",
		"resolved": "https://npm.intra.longguikeji.com/once/-/once-1.4.0.tgz",
		"integrity": "sha1-WDsap3WWHUsROsF9nFC6753Xa9E=",
		"dev": true,
		"requires": {
		  "wrappy": "1"
		}
	  },
	  "os-browserify": {
		"version": "0.3.0",
		"resolved": "https://npm.intra.longguikeji.com/os-browserify/-/os-browserify-0.3.0.tgz",
		"integrity": "sha1-hUNzx/XCMVkU/Jv8a9gjj92h7Cc=",
		"dev": true
	  },
	  "os-locale": {
		"version": "3.1.0",
		"resolved": "https://npm.intra.longguikeji.com/os-locale/-/os-locale-3.1.0.tgz",
		"integrity": "sha512-Z8l3R4wYWM40/52Z+S265okfFj8Kt2cC2MKY+xNi3kFs+XGI7WXu/I309QQQYbRW4ijiZ+yxs9pqEhJh0DqW3Q==",
		"dev": true,
		"requires": {
		  "execa": "^1.0.0",
		  "lcid": "^2.0.0",
		  "mem": "^4.0.0"
		}
	  },
	  "p-defer": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/p-defer/-/p-defer-1.0.0.tgz",
		"integrity": "sha1-n26xgvbJqozXQwBKfU+WsZaw+ww=",
		"dev": true
	  },
	  "p-finally": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/p-finally/-/p-finally-1.0.0.tgz",
		"integrity": "sha1-P7z7FbiZpEEjs0ttzBi3JDNqLK4=",
		"dev": true
	  },
	  "p-is-promise": {
		"version": "2.1.0",
		"resolved": "https://npm.intra.longguikeji.com/p-is-promise/-/p-is-promise-2.1.0.tgz",
		"integrity": "sha512-Y3W0wlRPK8ZMRbNq97l4M5otioeA5lm1z7bkNkxCka8HSPjR0xRWmpCmc9utiaLP9Jb1eD8BgeIxTW4AIF45Pg==",
		"dev": true
	  },
	  "p-limit": {
		"version": "2.2.2",
		"resolved": "https://npm.intra.longguikeji.com/p-limit/-/p-limit-2.2.2.tgz",
		"integrity": "sha512-WGR+xHecKTr7EbUEhyLSh5Dube9JtdiG78ufaeLxTgpudf/20KqyMioIUZJAezlTIi6evxuoUs9YXc11cU+yzQ==",
		"requires": {
		  "p-try": "^2.0.0"
		}
	  },
	  "p-locate": {
		"version": "4.1.0",
		"resolved": "https://npm.intra.longguikeji.com/p-locate/-/p-locate-4.1.0.tgz",
		"integrity": "sha512-R79ZZ/0wAxKGu3oYMlz8jy/kbhsNrS7SKZ7PxEHBgJ5+F2mtFW2fK2cOtBh1cHYkQsbzFV7I+EoRKe6Yt0oK7A==",
		"requires": {
		  "p-limit": "^2.2.0"
		}
	  },
	  "p-try": {
		"version": "2.2.0",
		"resolved": "https://npm.intra.longguikeji.com/p-try/-/p-try-2.2.0.tgz",
		"integrity": "sha512-R4nPAVTAU0B9D35/Gk3uJf/7XYbQcyohSKdvAxIRSNghFl4e71hVoGnBNQz9cWaXxO2I10KTC+3jMdvvoKw6dQ=="
	  },
	  "pako": {
		"version": "1.0.10",
		"resolved": "https://npm.intra.longguikeji.com/pako/-/pako-1.0.10.tgz",
		"integrity": "sha512-0DTvPVU3ed8+HNXOu5Bs+o//Mbdj9VNQMUOe9oKCwh8l0GNwpTDMKCWbRjgtD291AWnkAgkqA/LOnQS8AmS1tw==",
		"dev": true
	  },
	  "parallel-transform": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/parallel-transform/-/parallel-transform-1.2.0.tgz",
		"integrity": "sha512-P2vSmIu38uIlvdcU7fDkyrxj33gTUy/ABO5ZUbGowxNCopBq/OoD42bP4UmMrJoPyk4Uqf0mu3mtWBhHCZD8yg==",
		"dev": true,
		"requires": {
		  "cyclist": "^1.0.1",
		  "inherits": "^2.0.3",
		  "readable-stream": "^2.1.5"
		}
	  },
	  "parse-asn1": {
		"version": "5.1.5",
		"resolved": "https://npm.intra.longguikeji.com/parse-asn1/-/parse-asn1-5.1.5.tgz",
		"integrity": "sha512-jkMYn1dcJqF6d5CpU689bq7w/b5ALS9ROVSpQDPrZsqqesUJii9qutvoT5ltGedNXMO2e16YUWIghG9KxaViTQ==",
		"dev": true,
		"requires": {
		  "asn1.js": "^4.0.0",
		  "browserify-aes": "^1.0.0",
		  "create-hash": "^1.1.0",
		  "evp_bytestokey": "^1.0.0",
		  "pbkdf2": "^3.0.3",
		  "safe-buffer": "^5.1.1"
		}
	  },
	  "parse-passwd": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/parse-passwd/-/parse-passwd-1.0.0.tgz",
		"integrity": "sha1-bVuTSkVpk7I9N/QKOC1vFmao5cY=",
		"dev": true
	  },
	  "parseurl": {
		"version": "1.3.3",
		"resolved": "https://npm.intra.longguikeji.com/parseurl/-/parseurl-1.3.3.tgz",
		"integrity": "sha512-CiyeOxFT/JZyN5m0z9PfXw4SCBJ6Sygz1Dpl0wqjlhDEGGBP1GnsUVEL0p63hoG1fcj3fHynXi9NYO4nWOL+qQ=="
	  },
	  "pascalcase": {
		"version": "0.1.1",
		"resolved": "https://npm.intra.longguikeji.com/pascalcase/-/pascalcase-0.1.1.tgz",
		"integrity": "sha1-s2PlXoAGym/iF4TS2yK9FdeRfxQ=",
		"dev": true
	  },
	  "path-browserify": {
		"version": "0.0.1",
		"resolved": "https://npm.intra.longguikeji.com/path-browserify/-/path-browserify-0.0.1.tgz",
		"integrity": "sha512-BapA40NHICOS+USX9SN4tyhq+A2RrN/Ws5F0Z5aMHDp98Fl86lX8Oti8B7uN93L4Ifv4fHOEA+pQw87gmMO/lQ==",
		"dev": true
	  },
	  "path-dirname": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/path-dirname/-/path-dirname-1.0.2.tgz",
		"integrity": "sha1-zDPSTVJeCZpTiMAzbG4yuRYGCeA=",
		"dev": true
	  },
	  "path-exists": {
		"version": "4.0.0",
		"resolved": "https://npm.intra.longguikeji.com/path-exists/-/path-exists-4.0.0.tgz",
		"integrity": "sha512-ak9Qy5Q7jYb2Wwcey5Fpvg2KoAc/ZIhLSLOSBmRmygPsGwkVVt0fZa0qrtMz+m6tJTAHfZQ8FnmB4MG4LWy7/w=="
	  },
	  "path-is-absolute": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/path-is-absolute/-/path-is-absolute-1.0.1.tgz",
		"integrity": "sha1-F0uSaHNVNP+8es5r9TpanhtcX18=",
		"dev": true
	  },
	  "path-key": {
		"version": "2.0.1",
		"resolved": "https://npm.intra.longguikeji.com/path-key/-/path-key-2.0.1.tgz",
		"integrity": "sha1-QRyttXTFoUDTpLGRDUDYDMn0C0A=",
		"dev": true
	  },
	  "path-parse": {
		"version": "1.0.6",
		"resolved": "https://npm.intra.longguikeji.com/path-parse/-/path-parse-1.0.6.tgz",
		"integrity": "sha512-GSmOT2EbHrINBf9SR7CDELwlJ8AENk3Qn7OikK4nFYAu3Ote2+JYNVvkpAEQm3/TLNEJFD/xZJjzyxg3KBWOzw==",
		"dev": true
	  },
	  "path-to-regexp": {
		"version": "0.1.7",
		"resolved": "https://npm.intra.longguikeji.com/path-to-regexp/-/path-to-regexp-0.1.7.tgz",
		"integrity": "sha1-32BBeABfUi8V60SQ5yR6G/qmf4w="
	  },
	  "pbkdf2": {
		"version": "3.0.17",
		"resolved": "https://npm.intra.longguikeji.com/pbkdf2/-/pbkdf2-3.0.17.tgz",
		"integrity": "sha512-U/il5MsrZp7mGg3mSQfn742na2T+1/vHDCG5/iTI3X9MKUuYUZVLQhyRsg06mCgDBTd57TxzgZt7P+fYfjRLtA==",
		"dev": true,
		"requires": {
		  "create-hash": "^1.1.2",
		  "create-hmac": "^1.1.4",
		  "ripemd160": "^2.0.1",
		  "safe-buffer": "^5.0.1",
		  "sha.js": "^2.4.8"
		}
	  },
	  "pify": {
		"version": "4.0.1",
		"resolved": "https://npm.intra.longguikeji.com/pify/-/pify-4.0.1.tgz",
		"integrity": "sha512-uB80kBFb/tfd68bVleG9T5GGsGPjJrLAUpR5PZIrhBnIaRTQRjqdJSsIKkOP6OAIFbj7GOrcudc5pNjZ+geV2g==",
		"dev": true
	  },
	  "pkg-dir": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/pkg-dir/-/pkg-dir-3.0.0.tgz",
		"integrity": "sha512-/E57AYkoeQ25qkxMj5PBOVgF8Kiu/h7cYS30Z5+R7WaiCCBfLq58ZI/dSeaEKb9WVJV5n/03QwrN3IeWIFllvw==",
		"dev": true,
		"requires": {
		  "find-up": "^3.0.0"
		},
		"dependencies": {
		  "find-up": {
			"version": "3.0.0",
			"resolved": "https://npm.intra.longguikeji.com/find-up/-/find-up-3.0.0.tgz",
			"integrity": "sha512-1yD6RmLI1XBfxugvORwlck6f75tYL+iR0jqwsOrOxMZyGYqUuDhJ0l4AXdO1iX/FTs9cBAMEk1gWSEx1kSbylg==",
			"dev": true,
			"requires": {
			  "locate-path": "^3.0.0"
			}
		  },
		  "locate-path": {
			"version": "3.0.0",
			"resolved": "https://npm.intra.longguikeji.com/locate-path/-/locate-path-3.0.0.tgz",
			"integrity": "sha512-7AO748wWnIhNqAuaty2ZWHkQHRSNfPVIsPIfwEOWO22AmaoVrWavlOcMR5nzTLNYvp36X220/maaRsrec1G65A==",
			"dev": true,
			"requires": {
			  "p-locate": "^3.0.0",
			  "path-exists": "^3.0.0"
			}
		  },
		  "p-locate": {
			"version": "3.0.0",
			"resolved": "https://npm.intra.longguikeji.com/p-locate/-/p-locate-3.0.0.tgz",
			"integrity": "sha512-x+12w/To+4GFfgJhBEpiDcLozRJGegY+Ei7/z0tSLkMmxGZNybVMSfWj9aJn8Z5Fc7dBUNJOOVgPv2H7IwulSQ==",
			"dev": true,
			"requires": {
			  "p-limit": "^2.0.0"
			}
		  },
		  "path-exists": {
			"version": "3.0.0",
			"resolved": "https://npm.intra.longguikeji.com/path-exists/-/path-exists-3.0.0.tgz",
			"integrity": "sha1-zg6+ql94yxiSXqfYENe1mwEP1RU=",
			"dev": true
		  }
		}
	  },
	  "posix-character-classes": {
		"version": "0.1.1",
		"resolved": "https://npm.intra.longguikeji.com/posix-character-classes/-/posix-character-classes-0.1.1.tgz",
		"integrity": "sha1-AerA/jta9xoqbAL+q7jB/vfgDqs=",
		"dev": true
	  },
	  "prettier": {
		"version": "1.19.1",
		"resolved": "https://npm.intra.longguikeji.com/prettier/-/prettier-1.19.1.tgz",
		"integrity": "sha512-s7PoyDv/II1ObgQunCbB9PdLmUcBZcnWOcxDh7O0N/UwDEsHyqkW+Qh28jW+mVuCdx7gLB0BotYI1Y6uI9iyew==",
		"dev": true
	  },
	  "private": {
		"version": "0.1.8",
		"resolved": "https://npm.intra.longguikeji.com/private/-/private-0.1.8.tgz",
		"integrity": "sha512-VvivMrbvd2nKkiG38qjULzlc+4Vx4wm/whI9pQD35YrARNnhxeiRktSOhSukRLFNlzg6Br/cJPet5J/u19r/mg==",
		"dev": true
	  },
	  "process": {
		"version": "0.11.10",
		"resolved": "https://npm.intra.longguikeji.com/process/-/process-0.11.10.tgz",
		"integrity": "sha1-czIwDoQBYb2j5podHZGn1LwW8YI=",
		"dev": true
	  },
	  "process-nextick-args": {
		"version": "2.0.1",
		"resolved": "https://npm.intra.longguikeji.com/process-nextick-args/-/process-nextick-args-2.0.1.tgz",
		"integrity": "sha512-3ouUOpQhtgrbOa17J7+uxOTpITYWaGP7/AhoR3+A+/1e9skrzelGi/dXzEYyvbxubEF6Wn2ypscTKiKJFFn1ag==",
		"dev": true
	  },
	  "promise-inflight": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/promise-inflight/-/promise-inflight-1.0.1.tgz",
		"integrity": "sha1-mEcocL8igTL8vdhoEputEsPAKeM=",
		"dev": true
	  },
	  "proxy-addr": {
		"version": "2.0.5",
		"resolved": "https://npm.intra.longguikeji.com/proxy-addr/-/proxy-addr-2.0.5.tgz",
		"integrity": "sha512-t/7RxHXPH6cJtP0pRG6smSr9QJidhB+3kXu0KgXnbGYMgzEnUxRQ4/LDdfOwZEMyIh3/xHb8PX3t+lfL9z+YVQ==",
		"requires": {
		  "forwarded": "~0.1.2",
		  "ipaddr.js": "1.9.0"
		}
	  },
	  "prr": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/prr/-/prr-1.0.1.tgz",
		"integrity": "sha1-0/wRS6BplaRexok/SEzrHXj19HY=",
		"dev": true
	  },
	  "public-encrypt": {
		"version": "4.0.3",
		"resolved": "https://npm.intra.longguikeji.com/public-encrypt/-/public-encrypt-4.0.3.tgz",
		"integrity": "sha512-zVpa8oKZSz5bTMTFClc1fQOnyyEzpl5ozpi1B5YcvBrdohMjH2rfsBtyXcuNuwjsDIXmBYlF2N5FlJYhR29t8Q==",
		"dev": true,
		"requires": {
		  "bn.js": "^4.1.0",
		  "browserify-rsa": "^4.0.0",
		  "create-hash": "^1.1.0",
		  "parse-asn1": "^5.0.0",
		  "randombytes": "^2.0.1",
		  "safe-buffer": "^5.1.2"
		}
	  },
	  "pump": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/pump/-/pump-3.0.0.tgz",
		"integrity": "sha512-LwZy+p3SFs1Pytd/jYct4wpv49HiYCqd9Rlc5ZVdk0V+8Yzv6jR5Blk3TRmPL1ft69TxP0IMZGJ+WPFU2BFhww==",
		"dev": true,
		"requires": {
		  "end-of-stream": "^1.1.0",
		  "once": "^1.3.1"
		}
	  },
	  "pumpify": {
		"version": "1.5.1",
		"resolved": "https://npm.intra.longguikeji.com/pumpify/-/pumpify-1.5.1.tgz",
		"integrity": "sha512-oClZI37HvuUJJxSKKrC17bZ9Cu0ZYhEAGPsPUy9KlMUmv9dKX2o77RUmq7f3XjIxbwyGwYzbzQ1L2Ks8sIradQ==",
		"dev": true,
		"requires": {
		  "duplexify": "^3.6.0",
		  "inherits": "^2.0.3",
		  "pump": "^2.0.0"
		},
		"dependencies": {
		  "pump": {
			"version": "2.0.1",
			"resolved": "https://npm.intra.longguikeji.com/pump/-/pump-2.0.1.tgz",
			"integrity": "sha512-ruPMNRkN3MHP1cWJc9OWr+T/xDP0jhXYCLfJcBuX54hhfIBnaQmAUMfDcG4DM5UMWByBbJY69QSphm3jtDKIkA==",
			"dev": true,
			"requires": {
			  "end-of-stream": "^1.1.0",
			  "once": "^1.3.1"
			}
		  }
		}
	  },
	  "punycode": {
		"version": "2.1.1",
		"resolved": "https://npm.intra.longguikeji.com/punycode/-/punycode-2.1.1.tgz",
		"integrity": "sha512-XRsRjdf+j5ml+y/6GKHPZbrF/8p2Yga0JPtdqTIY2Xe5ohJPD9saDJJLPvp9+NSBprVvevdXZybnj2cv8OEd0A==",
		"dev": true
	  },
	  "qs": {
		"version": "6.7.0",
		"resolved": "https://npm.intra.longguikeji.com/qs/-/qs-6.7.0.tgz",
		"integrity": "sha512-VCdBRNFTX1fyE7Nb6FYoURo/SPe62QCaAyzJvUjwRaIsc+NePBEniHlvxFmmX56+HZphIGtV0XeCirBtpDrTyQ=="
	  },
	  "querystring": {
		"version": "0.2.0",
		"resolved": "https://npm.intra.longguikeji.com/querystring/-/querystring-0.2.0.tgz",
		"integrity": "sha1-sgmEkgO7Jd+CDadW50cAWHhSFiA=",
		"dev": true
	  },
	  "querystring-es3": {
		"version": "0.2.1",
		"resolved": "https://npm.intra.longguikeji.com/querystring-es3/-/querystring-es3-0.2.1.tgz",
		"integrity": "sha1-nsYfeQSYdXB9aUFFlv2Qek1xHnM=",
		"dev": true
	  },
	  "randombytes": {
		"version": "2.1.0",
		"resolved": "https://npm.intra.longguikeji.com/randombytes/-/randombytes-2.1.0.tgz",
		"integrity": "sha512-vYl3iOX+4CKUWuxGi9Ukhie6fsqXqS9FE2Zaic4tNFD2N2QQaXOMFbuKK4QmDHC0JO6B1Zp41J0LpT0oR68amQ==",
		"dev": true,
		"requires": {
		  "safe-buffer": "^5.1.0"
		}
	  },
	  "randomfill": {
		"version": "1.0.4",
		"resolved": "https://npm.intra.longguikeji.com/randomfill/-/randomfill-1.0.4.tgz",
		"integrity": "sha512-87lcbR8+MhcWcUiQ+9e+Rwx8MyR2P7qnt15ynUlbm3TU/fjbgz4GsvfSUDTemtCCtVCqb4ZcEFlyPNTh9bBTLw==",
		"dev": true,
		"requires": {
		  "randombytes": "^2.0.5",
		  "safe-buffer": "^5.1.0"
		}
	  },
	  "range-parser": {
		"version": "1.2.1",
		"resolved": "https://npm.intra.longguikeji.com/range-parser/-/range-parser-1.2.1.tgz",
		"integrity": "sha512-Hrgsx+orqoygnmhFbKaHE6c296J+HTAQXoxEF6gNupROmmGJRoyzfG3ccAveqCBrwr/2yxQ5BVd/GTl5agOwSg=="
	  },
	  "raw-body": {
		"version": "2.4.0",
		"resolved": "https://npm.intra.longguikeji.com/raw-body/-/raw-body-2.4.0.tgz",
		"integrity": "sha512-4Oz8DUIwdvoa5qMJelxipzi/iJIi40O5cGV1wNYp5hvZP8ZN0T+jiNkL0QepXs+EsQ9XJ8ipEDoiH70ySUJP3Q==",
		"requires": {
		  "bytes": "3.1.0",
		  "http-errors": "1.7.2",
		  "iconv-lite": "0.4.24",
		  "unpipe": "1.0.0"
		}
	  },
	  "readable-stream": {
		"version": "2.3.7",
		"resolved": "https://npm.intra.longguikeji.com/readable-stream/-/readable-stream-2.3.7.tgz",
		"integrity": "sha512-Ebho8K4jIbHAxnuxi7o42OrZgF/ZTNcsZj6nRKyUmkhLFq8CHItp/fy6hQZuZmP/n3yZ9VBUbp4zz/mX8hmYPw==",
		"dev": true,
		"requires": {
		  "core-util-is": "~1.0.0",
		  "inherits": "~2.0.3",
		  "isarray": "~1.0.0",
		  "process-nextick-args": "~2.0.0",
		  "safe-buffer": "~5.1.1",
		  "string_decoder": "~1.1.1",
		  "util-deprecate": "~1.0.1"
		}
	  },
	  "readdirp": {
		"version": "2.2.1",
		"resolved": "https://npm.intra.longguikeji.com/readdirp/-/readdirp-2.2.1.tgz",
		"integrity": "sha512-1JU/8q+VgFZyxwrJ+SVIOsh+KywWGpds3NTqikiKpDMZWScmAYyKIgqkO+ARvNWJfXeXR1zxz7aHF4u4CyH6vQ==",
		"dev": true,
		"requires": {
		  "graceful-fs": "^4.1.11",
		  "micromatch": "^3.1.10",
		  "readable-stream": "^2.0.2"
		}
	  },
	  "regenerate": {
		"version": "1.4.0",
		"resolved": "https://npm.intra.longguikeji.com/regenerate/-/regenerate-1.4.0.tgz",
		"integrity": "sha512-1G6jJVDWrt0rK99kBjvEtziZNCICAuvIPkSiUFIQxVP06RCVpq3dmDo2oi6ABpYaDYaTRr67BEhL8r1wgEZZKg==",
		"dev": true
	  },
	  "regenerate-unicode-properties": {
		"version": "8.1.0",
		"resolved": "https://npm.intra.longguikeji.com/regenerate-unicode-properties/-/regenerate-unicode-properties-8.1.0.tgz",
		"integrity": "sha512-LGZzkgtLY79GeXLm8Dp0BVLdQlWICzBnJz/ipWUgo59qBaZ+BHtq51P2q1uVZlppMuUAT37SDk39qUbjTWB7bA==",
		"dev": true,
		"requires": {
		  "regenerate": "^1.4.0"
		}
	  },
	  "regenerator-transform": {
		"version": "0.14.1",
		"resolved": "https://npm.intra.longguikeji.com/regenerator-transform/-/regenerator-transform-0.14.1.tgz",
		"integrity": "sha512-flVuee02C3FKRISbxhXl9mGzdbWUVHubl1SMaknjxkFB1/iqpJhArQUvRxOOPEc/9tAiX0BaQ28FJH10E4isSQ==",
		"dev": true,
		"requires": {
		  "private": "^0.1.6"
		}
	  },
	  "regex-not": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/regex-not/-/regex-not-1.0.2.tgz",
		"integrity": "sha512-J6SDjUgDxQj5NusnOtdFxDwN/+HWykR8GELwctJ7mdqhcyy1xEc4SRFHUXvxTp661YaVKAjfRLZ9cCqS6tn32A==",
		"dev": true,
		"requires": {
		  "extend-shallow": "^3.0.2",
		  "safe-regex": "^1.1.0"
		}
	  },
	  "regexpu-core": {
		"version": "4.6.0",
		"resolved": "https://npm.intra.longguikeji.com/regexpu-core/-/regexpu-core-4.6.0.tgz",
		"integrity": "sha512-YlVaefl8P5BnFYOITTNzDvan1ulLOiXJzCNZxduTIosN17b87h3bvG9yHMoHaRuo88H4mQ06Aodj5VtYGGGiTg==",
		"dev": true,
		"requires": {
		  "regenerate": "^1.4.0",
		  "regenerate-unicode-properties": "^8.1.0",
		  "regjsgen": "^0.5.0",
		  "regjsparser": "^0.6.0",
		  "unicode-match-property-ecmascript": "^1.0.4",
		  "unicode-match-property-value-ecmascript": "^1.1.0"
		}
	  },
	  "regjsgen": {
		"version": "0.5.1",
		"resolved": "https://npm.intra.longguikeji.com/regjsgen/-/regjsgen-0.5.1.tgz",
		"integrity": "sha512-5qxzGZjDs9w4tzT3TPhCJqWdCc3RLYwy9J2NB0nm5Lz+S273lvWcpjaTGHsT1dc6Hhfq41uSEOw8wBmxrKOuyg==",
		"dev": true
	  },
	  "regjsparser": {
		"version": "0.6.2",
		"resolved": "https://npm.intra.longguikeji.com/regjsparser/-/regjsparser-0.6.2.tgz",
		"integrity": "sha512-E9ghzUtoLwDekPT0DYCp+c4h+bvuUpe6rRHCTYn6eGoqj1LgKXxT6I0Il4WbjhQkOghzi/V+y03bPKvbllL93Q==",
		"dev": true,
		"requires": {
		  "jsesc": "~0.5.0"
		},
		"dependencies": {
		  "jsesc": {
			"version": "0.5.0",
			"resolved": "https://npm.intra.longguikeji.com/jsesc/-/jsesc-0.5.0.tgz",
			"integrity": "sha1-597mbjXW/Bb3EP6R1c9p9w8IkR0=",
			"dev": true
		  }
		}
	  },
	  "remove-trailing-separator": {
		"version": "1.1.0",
		"resolved": "https://npm.intra.longguikeji.com/remove-trailing-separator/-/remove-trailing-separator-1.1.0.tgz",
		"integrity": "sha1-wkvOKig62tW8P1jg1IJJuSN52O8=",
		"dev": true
	  },
	  "repeat-element": {
		"version": "1.1.3",
		"resolved": "https://npm.intra.longguikeji.com/repeat-element/-/repeat-element-1.1.3.tgz",
		"integrity": "sha512-ahGq0ZnV5m5XtZLMb+vP76kcAM5nkLqk0lpqAuojSKGgQtn4eRi4ZZGm2olo2zKFH+sMsWaqOCW1dqAnOru72g==",
		"dev": true
	  },
	  "repeat-string": {
		"version": "1.6.1",
		"resolved": "https://npm.intra.longguikeji.com/repeat-string/-/repeat-string-1.6.1.tgz",
		"integrity": "sha1-jcrkcOHIirwtYA//Sndihtp15jc=",
		"dev": true
	  },
	  "require-directory": {
		"version": "2.1.1",
		"resolved": "https://npm.intra.longguikeji.com/require-directory/-/require-directory-2.1.1.tgz",
		"integrity": "sha1-jGStX9MNqxyXbiNE/+f3kqam30I="
	  },
	  "require-main-filename": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/require-main-filename/-/require-main-filename-2.0.0.tgz",
		"integrity": "sha512-NKN5kMDylKuldxYLSUfrbo5Tuzh4hd+2E8NPPX02mZtn1VuREQToYe/ZdlJy+J3uCpfaiGF05e7B8W0iXbQHmg=="
	  },
	  "resolve": {
		"version": "1.14.2",
		"resolved": "https://npm.intra.longguikeji.com/resolve/-/resolve-1.14.2.tgz",
		"integrity": "sha512-EjlOBLBO1kxsUxsKjLt7TAECyKW6fOh1VRkykQkKGzcBbjjPIxBqGh0jf7GJ3k/f5mxMqW3htMD3WdTUVtW8HQ==",
		"dev": true,
		"requires": {
		  "path-parse": "^1.0.6"
		}
	  },
	  "resolve-cwd": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/resolve-cwd/-/resolve-cwd-2.0.0.tgz",
		"integrity": "sha1-AKn3OHVW4nA46uIyyqNypqWbZlo=",
		"dev": true,
		"requires": {
		  "resolve-from": "^3.0.0"
		}
	  },
	  "resolve-dir": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/resolve-dir/-/resolve-dir-1.0.1.tgz",
		"integrity": "sha1-eaQGRMNivoLybv/nOcm7U4IEb0M=",
		"dev": true,
		"requires": {
		  "expand-tilde": "^2.0.0",
		  "global-modules": "^1.0.0"
		},
		"dependencies": {
		  "global-modules": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/global-modules/-/global-modules-1.0.0.tgz",
			"integrity": "sha512-sKzpEkf11GpOFuw0Zzjzmt4B4UZwjOcG757PPvrfhxcLFbq0wpsgpOqxpxtxFiCG4DtG93M6XRVbF2oGdev7bg==",
			"dev": true,
			"requires": {
			  "global-prefix": "^1.0.1",
			  "is-windows": "^1.0.1",
			  "resolve-dir": "^1.0.0"
			}
		  }
		}
	  },
	  "resolve-from": {
		"version": "3.0.0",
		"resolved": "https://npm.intra.longguikeji.com/resolve-from/-/resolve-from-3.0.0.tgz",
		"integrity": "sha1-six699nWiBvItuZTM17rywoYh0g=",
		"dev": true
	  },
	  "resolve-url": {
		"version": "0.2.1",
		"resolved": "https://npm.intra.longguikeji.com/resolve-url/-/resolve-url-0.2.1.tgz",
		"integrity": "sha1-LGN/53yJOv0qZj/iGqkIAGjiBSo=",
		"dev": true
	  },
	  "ret": {
		"version": "0.1.15",
		"resolved": "https://npm.intra.longguikeji.com/ret/-/ret-0.1.15.tgz",
		"integrity": "sha512-TTlYpa+OL+vMMNG24xSlQGEJ3B/RzEfUlLct7b5G/ytav+wPrplCpVMFuwzXbkecJrb6IYo1iFb0S9v37754mg==",
		"dev": true
	  },
	  "rimraf": {
		"version": "2.7.1",
		"resolved": "https://npm.intra.longguikeji.com/rimraf/-/rimraf-2.7.1.tgz",
		"integrity": "sha512-uWjbaKIK3T1OSVptzX7Nl6PvQ3qAGtKEtVRjRuazjfL3Bx5eI409VZSqgND+4UNnmzLVdPj9FqFJNPqBZFve4w==",
		"dev": true,
		"requires": {
		  "glob": "^7.1.3"
		}
	  },
	  "ripemd160": {
		"version": "2.0.2",
		"resolved": "https://npm.intra.longguikeji.com/ripemd160/-/ripemd160-2.0.2.tgz",
		"integrity": "sha512-ii4iagi25WusVoiC4B4lq7pbXfAp3D9v5CwfkY33vffw2+pkDjY1D8GaN7spsxvCSx8dkPqOZCEZyfxcmJG2IA==",
		"dev": true,
		"requires": {
		  "hash-base": "^3.0.0",
		  "inherits": "^2.0.1"
		}
	  },
	  "run-queue": {
		"version": "1.0.3",
		"resolved": "https://npm.intra.longguikeji.com/run-queue/-/run-queue-1.0.3.tgz",
		"integrity": "sha1-6Eg5bwV9Ij8kOGkkYY4laUFh7Ec=",
		"dev": true,
		"requires": {
		  "aproba": "^1.1.1"
		}
	  },
	  "safe-buffer": {
		"version": "5.1.2",
		"resolved": "https://npm.intra.longguikeji.com/safe-buffer/-/safe-buffer-5.1.2.tgz",
		"integrity": "sha512-Gd2UZBJDkXlY7GbJxfsE8/nvKkUEU1G38c1siN6QP6a9PT9MmHB8GnpscSmMJSoF8LOIrt8ud/wPtojys4G6+g=="
	  },
	  "safe-regex": {
		"version": "1.1.0",
		"resolved": "https://npm.intra.longguikeji.com/safe-regex/-/safe-regex-1.1.0.tgz",
		"integrity": "sha1-QKNmnzsHfR6UPURinhV91IAjvy4=",
		"dev": true,
		"requires": {
		  "ret": "~0.1.10"
		}
	  },
	  "safer-buffer": {
		"version": "2.1.2",
		"resolved": "https://npm.intra.longguikeji.com/safer-buffer/-/safer-buffer-2.1.2.tgz",
		"integrity": "sha512-YZo3K82SD7Riyi0E1EQPojLz7kpepnSQI9IyPbHHg1XXXevb5dJI7tpyN2ADxGcQbHG7vcyRHk0cbwqcQriUtg=="
	  },
	  "schema-utils": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/schema-utils/-/schema-utils-1.0.0.tgz",
		"integrity": "sha512-i27Mic4KovM/lnGsy8whRCHhc7VicJajAjTrYg11K9zfZXnYIt4k5F+kZkwjnrhKzLic/HLU4j11mjsz2G/75g==",
		"dev": true,
		"requires": {
		  "ajv": "^6.1.0",
		  "ajv-errors": "^1.0.0",
		  "ajv-keywords": "^3.1.0"
		}
	  },
	  "semver": {
		"version": "5.7.1",
		"resolved": "https://npm.intra.longguikeji.com/semver/-/semver-5.7.1.tgz",
		"integrity": "sha512-sauaDf/PZdVgrLTNYHRtpXa1iRiKcaebiKQ1BJdpQlWH2lCvexQdX55snPFyK7QzpudqbCI0qXFfOasHdyNDGQ==",
		"dev": true
	  },
	  "send": {
		"version": "0.17.1",
		"resolved": "https://npm.intra.longguikeji.com/send/-/send-0.17.1.tgz",
		"integrity": "sha512-BsVKsiGcQMFwT8UxypobUKyv7irCNRHk1T0G680vk88yf6LBByGcZJOTJCrTP2xVN6yI+XjPJcNuE3V4fT9sAg==",
		"requires": {
		  "debug": "2.6.9",
		  "depd": "~1.1.2",
		  "destroy": "~1.0.4",
		  "encodeurl": "~1.0.2",
		  "escape-html": "~1.0.3",
		  "etag": "~1.8.1",
		  "fresh": "0.5.2",
		  "http-errors": "~1.7.2",
		  "mime": "1.6.0",
		  "ms": "2.1.1",
		  "on-finished": "~2.3.0",
		  "range-parser": "~1.2.1",
		  "statuses": "~1.5.0"
		},
		"dependencies": {
		  "debug": {
			"version": "2.6.9",
			"resolved": "https://npm.intra.longguikeji.com/debug/-/debug-2.6.9.tgz",
			"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
			"requires": {
			  "ms": "2.0.0"
			},
			"dependencies": {
			  "ms": {
				"version": "2.0.0",
				"resolved": "https://npm.intra.longguikeji.com/ms/-/ms-2.0.0.tgz",
				"integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g="
			  }
			}
		  },
		  "ms": {
			"version": "2.1.1",
			"resolved": "https://npm.intra.longguikeji.com/ms/-/ms-2.1.1.tgz",
			"integrity": "sha512-tgp+dl5cGk28utYktBsrFqA7HKgrhgPsg6Z/EfhWI4gl1Hwq8B/GmY/0oXZ6nF8hDVesS/FpnYaD/kOWhYQvyg=="
		  }
		}
	  },
	  "serialize-javascript": {
		"version": "2.1.2",
		"resolved": "https://npm.intra.longguikeji.com/serialize-javascript/-/serialize-javascript-2.1.2.tgz",
		"integrity": "sha512-rs9OggEUF0V4jUSecXazOYsLfu7OGK2qIn3c7IPBiffz32XniEp/TX9Xmc9LQfK2nQ2QKHvZ2oygKUGU0lG4jQ==",
		"dev": true
	  },
	  "serve-static": {
		"version": "1.14.1",
		"resolved": "https://npm.intra.longguikeji.com/serve-static/-/serve-static-1.14.1.tgz",
		"integrity": "sha512-JMrvUwE54emCYWlTI+hGrGv5I8dEwmco/00EvkzIIsR7MqrHonbD9pO2MOfFnpFntl7ecpZs+3mW+XbQZu9QCg==",
		"requires": {
		  "encodeurl": "~1.0.2",
		  "escape-html": "~1.0.3",
		  "parseurl": "~1.3.3",
		  "send": "0.17.1"
		}
	  },
	  "set-blocking": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/set-blocking/-/set-blocking-2.0.0.tgz",
		"integrity": "sha1-BF+XgtARrppoA93TgrJDkrPYkPc="
	  },
	  "set-value": {
		"version": "2.0.1",
		"resolved": "https://npm.intra.longguikeji.com/set-value/-/set-value-2.0.1.tgz",
		"integrity": "sha512-JxHc1weCN68wRY0fhCoXpyK55m/XPHafOmK4UWD7m2CI14GMcFypt4w/0+NV5f/ZMby2F6S2wwA7fgynh9gWSw==",
		"dev": true,
		"requires": {
		  "extend-shallow": "^2.0.1",
		  "is-extendable": "^0.1.1",
		  "is-plain-object": "^2.0.3",
		  "split-string": "^3.0.1"
		},
		"dependencies": {
		  "extend-shallow": {
			"version": "2.0.1",
			"resolved": "https://npm.intra.longguikeji.com/extend-shallow/-/extend-shallow-2.0.1.tgz",
			"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
			"dev": true,
			"requires": {
			  "is-extendable": "^0.1.0"
			}
		  }
		}
	  },
	  "setimmediate": {
		"version": "1.0.5",
		"resolved": "https://npm.intra.longguikeji.com/setimmediate/-/setimmediate-1.0.5.tgz",
		"integrity": "sha1-KQy7Iy4waULX1+qbg3Mqt4VvgoU=",
		"dev": true
	  },
	  "setprototypeof": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/setprototypeof/-/setprototypeof-1.1.1.tgz",
		"integrity": "sha512-JvdAWfbXeIGaZ9cILp38HntZSFSo3mWg6xGcJJsd+d4aRMOqauag1C63dJfDw7OaMYwEbHMOxEZ1lqVRYP2OAw=="
	  },
	  "sha.js": {
		"version": "2.4.11",
		"resolved": "https://npm.intra.longguikeji.com/sha.js/-/sha.js-2.4.11.tgz",
		"integrity": "sha512-QMEp5B7cftE7APOjk5Y6xgrbWu+WkLVQwk8JNjZ8nKRciZaByEW6MubieAiToS7+dwvrjGhH8jRXz3MVd0AYqQ==",
		"dev": true,
		"requires": {
		  "inherits": "^2.0.1",
		  "safe-buffer": "^5.0.1"
		}
	  },
	  "shebang-command": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/shebang-command/-/shebang-command-1.2.0.tgz",
		"integrity": "sha1-RKrGW2lbAzmJaMOfNj/uXer98eo=",
		"dev": true,
		"requires": {
		  "shebang-regex": "^1.0.0"
		}
	  },
	  "shebang-regex": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/shebang-regex/-/shebang-regex-1.0.0.tgz",
		"integrity": "sha1-2kL0l0DAtC2yypcoVxyxkMmO/qM=",
		"dev": true
	  },
	  "signal-exit": {
		"version": "3.0.2",
		"resolved": "https://npm.intra.longguikeji.com/signal-exit/-/signal-exit-3.0.2.tgz",
		"integrity": "sha1-tf3AjxKH6hF4Yo5BXiUTK3NkbG0=",
		"dev": true
	  },
	  "slash": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/slash/-/slash-2.0.0.tgz",
		"integrity": "sha512-ZYKh3Wh2z1PpEXWr0MpSBZ0V6mZHAQfYevttO11c51CaWjGTaadiKZ+wVt1PbMlDV5qhMFslpZCemhwOK7C89A==",
		"dev": true
	  },
	  "snapdragon": {
		"version": "0.8.2",
		"resolved": "https://npm.intra.longguikeji.com/snapdragon/-/snapdragon-0.8.2.tgz",
		"integrity": "sha512-FtyOnWN/wCHTVXOMwvSv26d+ko5vWlIDD6zoUJ7LW8vh+ZBC8QdljveRP+crNrtBwioEUWy/4dMtbBjA4ioNlg==",
		"dev": true,
		"requires": {
		  "base": "^0.11.1",
		  "debug": "^2.2.0",
		  "define-property": "^0.2.5",
		  "extend-shallow": "^2.0.1",
		  "map-cache": "^0.2.2",
		  "source-map": "^0.5.6",
		  "source-map-resolve": "^0.5.0",
		  "use": "^3.1.0"
		},
		"dependencies": {
		  "debug": {
			"version": "2.6.9",
			"resolved": "https://npm.intra.longguikeji.com/debug/-/debug-2.6.9.tgz",
			"integrity": "sha512-bC7ElrdJaJnPbAP+1EotYvqZsb3ecl5wi6Bfi6BJTUcNowp6cvspg0jXznRTKDjm/E7AdgFBVeAPVMNcKGsHMA==",
			"dev": true,
			"requires": {
			  "ms": "2.0.0"
			}
		  },
		  "define-property": {
			"version": "0.2.5",
			"resolved": "https://npm.intra.longguikeji.com/define-property/-/define-property-0.2.5.tgz",
			"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
			"dev": true,
			"requires": {
			  "is-descriptor": "^0.1.0"
			}
		  },
		  "extend-shallow": {
			"version": "2.0.1",
			"resolved": "https://npm.intra.longguikeji.com/extend-shallow/-/extend-shallow-2.0.1.tgz",
			"integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
			"dev": true,
			"requires": {
			  "is-extendable": "^0.1.0"
			}
		  }
		}
	  },
	  "snapdragon-node": {
		"version": "2.1.1",
		"resolved": "https://npm.intra.longguikeji.com/snapdragon-node/-/snapdragon-node-2.1.1.tgz",
		"integrity": "sha512-O27l4xaMYt/RSQ5TR3vpWCAB5Kb/czIcqUFOM/C4fYcLnbZUc1PkjTAMjof2pBWaSTwOUd6qUHcFGVGj7aIwnw==",
		"dev": true,
		"requires": {
		  "define-property": "^1.0.0",
		  "isobject": "^3.0.0",
		  "snapdragon-util": "^3.0.1"
		},
		"dependencies": {
		  "define-property": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/define-property/-/define-property-1.0.0.tgz",
			"integrity": "sha1-dp66rz9KY6rTr56NMEybvnm/sOY=",
			"dev": true,
			"requires": {
			  "is-descriptor": "^1.0.0"
			}
		  },
		  "is-accessor-descriptor": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/is-accessor-descriptor/-/is-accessor-descriptor-1.0.0.tgz",
			"integrity": "sha512-m5hnHTkcVsPfqx3AKlyttIPb7J+XykHvJP2B9bZDjlhLIoEq4XoK64Vg7boZlVWYK6LUY94dYPEE7Lh0ZkZKcQ==",
			"dev": true,
			"requires": {
			  "kind-of": "^6.0.0"
			}
		  },
		  "is-data-descriptor": {
			"version": "1.0.0",
			"resolved": "https://npm.intra.longguikeji.com/is-data-descriptor/-/is-data-descriptor-1.0.0.tgz",
			"integrity": "sha512-jbRXy1FmtAoCjQkVmIVYwuuqDFUbaOeDjmed1tOGPrsMhtJA4rD9tkgA0F1qJ3gRFRXcHYVkdeaP50Q5rE/jLQ==",
			"dev": true,
			"requires": {
			  "kind-of": "^6.0.0"
			}
		  },
		  "is-descriptor": {
			"version": "1.0.2",
			"resolved": "https://npm.intra.longguikeji.com/is-descriptor/-/is-descriptor-1.0.2.tgz",
			"integrity": "sha512-2eis5WqQGV7peooDyLmNEPUrps9+SXX5c9pL3xEB+4e9HnGuDa7mB7kHxHw4CbqS9k1T2hOH3miL8n8WtiYVtg==",
			"dev": true,
			"requires": {
			  "is-accessor-descriptor": "^1.0.0",
			  "is-data-descriptor": "^1.0.0",
			  "kind-of": "^6.0.2"
			}
		  }
		}
	  },
	  "snapdragon-util": {
		"version": "3.0.1",
		"resolved": "https://npm.intra.longguikeji.com/snapdragon-util/-/snapdragon-util-3.0.1.tgz",
		"integrity": "sha512-mbKkMdQKsjX4BAL4bRYTj21edOf8cN7XHdYUJEe+Zn99hVEYcMvKPct1IqNe7+AZPirn8BCDOQBHQZknqmKlZQ==",
		"dev": true,
		"requires": {
		  "kind-of": "^3.2.0"
		},
		"dependencies": {
		  "is-buffer": {
			"version": "1.1.6",
			"resolved": "https://npm.intra.longguikeji.com/is-buffer/-/is-buffer-1.1.6.tgz",
			"integrity": "sha512-NcdALwpXkTm5Zvvbk7owOUSvVvBKDgKP5/ewfXEznmQFfs4ZRmanOeKBTjRVjka3QFoN6XJ+9F3USqfHqTaU5w==",
			"dev": true
		  },
		  "kind-of": {
			"version": "3.2.2",
			"resolved": "https://npm.intra.longguikeji.com/kind-of/-/kind-of-3.2.2.tgz",
			"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
			"dev": true,
			"requires": {
			  "is-buffer": "^1.1.5"
			}
		  }
		}
	  },
	  "source-list-map": {
		"version": "2.0.1",
		"resolved": "https://npm.intra.longguikeji.com/source-list-map/-/source-list-map-2.0.1.tgz",
		"integrity": "sha512-qnQ7gVMxGNxsiL4lEuJwe/To8UnK7fAnmbGEEH8RpLouuKbeEm0lhbQVFIrNSuB+G7tVrAlVsZgETT5nljf+Iw==",
		"dev": true
	  },
	  "source-map": {
		"version": "0.5.7",
		"resolved": "https://npm.intra.longguikeji.com/source-map/-/source-map-0.5.7.tgz",
		"integrity": "sha1-igOdLRAh0i0eoUyA2OpGi6LvP8w=",
		"dev": true
	  },
	  "source-map-resolve": {
		"version": "0.5.3",
		"resolved": "https://npm.intra.longguikeji.com/source-map-resolve/-/source-map-resolve-0.5.3.tgz",
		"integrity": "sha512-Htz+RnsXWk5+P2slx5Jh3Q66vhQj1Cllm0zvnaY98+NFx+Dv2CF/f5O/t8x+KaNdrdIAsruNzoh/KpialbqAnw==",
		"dev": true,
		"requires": {
		  "atob": "^2.1.2",
		  "decode-uri-component": "^0.2.0",
		  "resolve-url": "^0.2.1",
		  "source-map-url": "^0.4.0",
		  "urix": "^0.1.0"
		}
	  },
	  "source-map-support": {
		"version": "0.5.16",
		"resolved": "https://npm.intra.longguikeji.com/source-map-support/-/source-map-support-0.5.16.tgz",
		"integrity": "sha512-efyLRJDr68D9hBBNIPWFjhpFzURh+KJykQwvMyW5UiZzYwoF6l4YMMDIJJEyFWxWCqfyxLzz6tSfUFR+kXXsVQ==",
		"dev": true,
		"requires": {
		  "buffer-from": "^1.0.0",
		  "source-map": "^0.6.0"
		},
		"dependencies": {
		  "source-map": {
			"version": "0.6.1",
			"resolved": "https://npm.intra.longguikeji.com/source-map/-/source-map-0.6.1.tgz",
			"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
			"dev": true
		  }
		}
	  },
	  "source-map-url": {
		"version": "0.4.0",
		"resolved": "https://npm.intra.longguikeji.com/source-map-url/-/source-map-url-0.4.0.tgz",
		"integrity": "sha1-PpNdfd1zYxuXZZlW1VEo6HtQhKM=",
		"dev": true
	  },
	  "split-string": {
		"version": "3.1.0",
		"resolved": "https://npm.intra.longguikeji.com/split-string/-/split-string-3.1.0.tgz",
		"integrity": "sha512-NzNVhJDYpwceVVii8/Hu6DKfD2G+NrQHlS/V/qgv763EYudVwEcMQNxd2lh+0VrUByXN/oJkl5grOhYWvQUYiw==",
		"dev": true,
		"requires": {
		  "extend-shallow": "^3.0.0"
		}
	  },
	  "ssri": {
		"version": "6.0.1",
		"resolved": "https://npm.intra.longguikeji.com/ssri/-/ssri-6.0.1.tgz",
		"integrity": "sha512-3Wge10hNcT1Kur4PDFwEieXSCMCJs/7WvSACcrMYrNp+b8kDL1/0wJch5Ni2WrtwEa2IO8OsVfeKIciKCDx/QA==",
		"dev": true,
		"requires": {
		  "figgy-pudding": "^3.5.1"
		}
	  },
	  "static-extend": {
		"version": "0.1.2",
		"resolved": "https://npm.intra.longguikeji.com/static-extend/-/static-extend-0.1.2.tgz",
		"integrity": "sha1-YICcOcv/VTNyJv1eC1IPNB8ftcY=",
		"dev": true,
		"requires": {
		  "define-property": "^0.2.5",
		  "object-copy": "^0.1.0"
		},
		"dependencies": {
		  "define-property": {
			"version": "0.2.5",
			"resolved": "https://npm.intra.longguikeji.com/define-property/-/define-property-0.2.5.tgz",
			"integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
			"dev": true,
			"requires": {
			  "is-descriptor": "^0.1.0"
			}
		  }
		}
	  },
	  "statuses": {
		"version": "1.5.0",
		"resolved": "https://npm.intra.longguikeji.com/statuses/-/statuses-1.5.0.tgz",
		"integrity": "sha1-Fhx9rBd2Wf2YEfQ3cfqZOBR4Yow="
	  },
	  "stream-browserify": {
		"version": "2.0.2",
		"resolved": "https://npm.intra.longguikeji.com/stream-browserify/-/stream-browserify-2.0.2.tgz",
		"integrity": "sha512-nX6hmklHs/gr2FuxYDltq8fJA1GDlxKQCz8O/IM4atRqBH8OORmBNgfvW5gG10GT/qQ9u0CzIvr2X5Pkt6ntqg==",
		"dev": true,
		"requires": {
		  "inherits": "~2.0.1",
		  "readable-stream": "^2.0.2"
		}
	  },
	  "stream-each": {
		"version": "1.2.3",
		"resolved": "https://npm.intra.longguikeji.com/stream-each/-/stream-each-1.2.3.tgz",
		"integrity": "sha512-vlMC2f8I2u/bZGqkdfLQW/13Zihpej/7PmSiMQsbYddxuTsJp8vRe2x2FvVExZg7FaOds43ROAuFJwPR4MTZLw==",
		"dev": true,
		"requires": {
		  "end-of-stream": "^1.1.0",
		  "stream-shift": "^1.0.0"
		}
	  },
	  "stream-http": {
		"version": "2.8.3",
		"resolved": "https://npm.intra.longguikeji.com/stream-http/-/stream-http-2.8.3.tgz",
		"integrity": "sha512-+TSkfINHDo4J+ZobQLWiMouQYB+UVYFttRA94FpEzzJ7ZdqcL4uUUQ7WkdkI4DSozGmgBUE/a47L+38PenXhUw==",
		"dev": true,
		"requires": {
		  "builtin-status-codes": "^3.0.0",
		  "inherits": "^2.0.1",
		  "readable-stream": "^2.3.6",
		  "to-arraybuffer": "^1.0.0",
		  "xtend": "^4.0.0"
		}
	  },
	  "stream-shift": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/stream-shift/-/stream-shift-1.0.1.tgz",
		"integrity": "sha512-AiisoFqQ0vbGcZgQPY1cdP2I76glaVA/RauYR4G4thNFgkTqr90yXTo4LYX60Jl+sIlPNHHdGSwo01AvbKUSVQ==",
		"dev": true
	  },
	  "string-width": {
		"version": "4.2.0",
		"resolved": "https://npm.intra.longguikeji.com/string-width/-/string-width-4.2.0.tgz",
		"integrity": "sha512-zUz5JD+tgqtuDjMhwIg5uFVV3dtqZ9yQJlZVfq4I01/K5Paj5UHj7VyrQOJvzawSVlKpObApbfD0Ed6yJc+1eg==",
		"requires": {
		  "emoji-regex": "^8.0.0",
		  "is-fullwidth-code-point": "^3.0.0",
		  "strip-ansi": "^6.0.0"
		}
	  },
	  "string_decoder": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/string_decoder/-/string_decoder-1.1.1.tgz",
		"integrity": "sha512-n/ShnvDi6FHbbVfviro+WojiFzv+s8MPMHBczVePfUpDJLwoLT0ht1l4YwBCbi8pJAveEEdnkHyPyTP/mzRfwg==",
		"dev": true,
		"requires": {
		  "safe-buffer": "~5.1.0"
		}
	  },
	  "strip-ansi": {
		"version": "6.0.0",
		"resolved": "https://npm.intra.longguikeji.com/strip-ansi/-/strip-ansi-6.0.0.tgz",
		"integrity": "sha512-AuvKTrTfQNYNIctbR1K/YGTR1756GycPsg7b9bdV9Duqur4gv6aKqHXah67Z8ImS7WEz5QVcOtlfW2rZEugt6w==",
		"requires": {
		  "ansi-regex": "^5.0.0"
		}
	  },
	  "strip-eof": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/strip-eof/-/strip-eof-1.0.0.tgz",
		"integrity": "sha1-u0P/VZim6wXYm1n80SnJgzE2Br8=",
		"dev": true
	  },
	  "supports-color": {
		"version": "5.5.0",
		"resolved": "https://npm.intra.longguikeji.com/supports-color/-/supports-color-5.5.0.tgz",
		"integrity": "sha512-QjVjwdXIt408MIiAqCX4oUKsgU2EqAGzs2Ppkm4aQYbjm+ZEWEcW4SfFNTr4uMNZma0ey4f5lgLrkB0aX0QMow==",
		"dev": true,
		"requires": {
		  "has-flag": "^3.0.0"
		}
	  },
	  "tapable": {
		"version": "1.1.3",
		"resolved": "https://npm.intra.longguikeji.com/tapable/-/tapable-1.1.3.tgz",
		"integrity": "sha512-4WK/bYZmj8xLr+HUCODHGF1ZFzsYffasLUgEiMBY4fgtltdO6B4WJtlSbPaDTLpYTcGVwM2qLnFTICEcNxs3kA==",
		"dev": true
	  },
	  "terser": {
		"version": "4.6.1",
		"resolved": "https://npm.intra.longguikeji.com/terser/-/terser-4.6.1.tgz",
		"integrity": "sha512-w0f2OWFD7ka3zwetgVAhNMeyzEbj39ht2Tb0qKflw9PmW9Qbo5tjTh01QJLkhO9t9RDDQYvk+WXqpECI2C6i2A==",
		"dev": true,
		"requires": {
		  "commander": "^2.20.0",
		  "source-map": "~0.6.1",
		  "source-map-support": "~0.5.12"
		},
		"dependencies": {
		  "commander": {
			"version": "2.20.3",
			"resolved": "https://npm.intra.longguikeji.com/commander/-/commander-2.20.3.tgz",
			"integrity": "sha512-GpVkmM8vF2vQUkj2LvZmD35JxeJOLCwJ9cUkugyk2nuhbv3+mJvpLYYt+0+USMxE+oj+ey/lJEnhZw75x/OMcQ==",
			"dev": true
		  },
		  "source-map": {
			"version": "0.6.1",
			"resolved": "https://npm.intra.longguikeji.com/source-map/-/source-map-0.6.1.tgz",
			"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
			"dev": true
		  }
		}
	  },
	  "terser-webpack-plugin": {
		"version": "1.4.3",
		"resolved": "https://npm.intra.longguikeji.com/terser-webpack-plugin/-/terser-webpack-plugin-1.4.3.tgz",
		"integrity": "sha512-QMxecFz/gHQwteWwSo5nTc6UaICqN1bMedC5sMtUc7y3Ha3Q8y6ZO0iCR8pq4RJC8Hjf0FEPEHZqcMB/+DFCrA==",
		"dev": true,
		"requires": {
		  "cacache": "^12.0.2",
		  "find-cache-dir": "^2.1.0",
		  "is-wsl": "^1.1.0",
		  "schema-utils": "^1.0.0",
		  "serialize-javascript": "^2.1.2",
		  "source-map": "^0.6.1",
		  "terser": "^4.1.2",
		  "webpack-sources": "^1.4.0",
		  "worker-farm": "^1.7.0"
		},
		"dependencies": {
		  "source-map": {
			"version": "0.6.1",
			"resolved": "https://npm.intra.longguikeji.com/source-map/-/source-map-0.6.1.tgz",
			"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
			"dev": true
		  }
		}
	  },
	  "through2": {
		"version": "2.0.5",
		"resolved": "https://npm.intra.longguikeji.com/through2/-/through2-2.0.5.tgz",
		"integrity": "sha512-/mrRod8xqpA+IHSLyGCQ2s8SPHiCDEeQJSep1jqLYeEUClOFG2Qsh+4FU6G9VeqpZnGW/Su8LQGc4YKni5rYSQ==",
		"dev": true,
		"requires": {
		  "readable-stream": "~2.3.6",
		  "xtend": "~4.0.1"
		}
	  },
	  "timers-browserify": {
		"version": "2.0.11",
		"resolved": "https://npm.intra.longguikeji.com/timers-browserify/-/timers-browserify-2.0.11.tgz",
		"integrity": "sha512-60aV6sgJ5YEbzUdn9c8kYGIqOubPoUdqQCul3SBAsRCZ40s6Y5cMcrW4dt3/k/EsbLVJNl9n6Vz3fTc+k2GeKQ==",
		"dev": true,
		"requires": {
		  "setimmediate": "^1.0.4"
		}
	  },
	  "to-arraybuffer": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/to-arraybuffer/-/to-arraybuffer-1.0.1.tgz",
		"integrity": "sha1-fSKbH8xjfkZsoIEYCDanqr/4P0M=",
		"dev": true
	  },
	  "to-fast-properties": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/to-fast-properties/-/to-fast-properties-2.0.0.tgz",
		"integrity": "sha1-3F5pjL0HkmW8c+A3doGk5Og/YW4=",
		"dev": true
	  },
	  "to-object-path": {
		"version": "0.3.0",
		"resolved": "https://npm.intra.longguikeji.com/to-object-path/-/to-object-path-0.3.0.tgz",
		"integrity": "sha1-KXWIt7Dn4KwI4E5nL4XB9JmeF68=",
		"dev": true,
		"requires": {
		  "kind-of": "^3.0.2"
		},
		"dependencies": {
		  "is-buffer": {
			"version": "1.1.6",
			"resolved": "https://npm.intra.longguikeji.com/is-buffer/-/is-buffer-1.1.6.tgz",
			"integrity": "sha512-NcdALwpXkTm5Zvvbk7owOUSvVvBKDgKP5/ewfXEznmQFfs4ZRmanOeKBTjRVjka3QFoN6XJ+9F3USqfHqTaU5w==",
			"dev": true
		  },
		  "kind-of": {
			"version": "3.2.2",
			"resolved": "https://npm.intra.longguikeji.com/kind-of/-/kind-of-3.2.2.tgz",
			"integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
			"dev": true,
			"requires": {
			  "is-buffer": "^1.1.5"
			}
		  }
		}
	  },
	  "to-regex": {
		"version": "3.0.2",
		"resolved": "https://npm.intra.longguikeji.com/to-regex/-/to-regex-3.0.2.tgz",
		"integrity": "sha512-FWtleNAtZ/Ki2qtqej2CXTOayOH9bHDQF+Q48VpWyDXjbYxA4Yz8iDB31zXOBUlOHHKidDbqGVrTUvQMPmBGBw==",
		"dev": true,
		"requires": {
		  "define-property": "^2.0.2",
		  "extend-shallow": "^3.0.2",
		  "regex-not": "^1.0.2",
		  "safe-regex": "^1.1.0"
		}
	  },
	  "to-regex-range": {
		"version": "2.1.1",
		"resolved": "https://npm.intra.longguikeji.com/to-regex-range/-/to-regex-range-2.1.1.tgz",
		"integrity": "sha1-fIDBe53+vlmeJzZ+DU3VWQFB2zg=",
		"dev": true,
		"requires": {
		  "is-number": "^3.0.0",
		  "repeat-string": "^1.6.1"
		}
	  },
	  "toidentifier": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/toidentifier/-/toidentifier-1.0.0.tgz",
		"integrity": "sha512-yaOH/Pk/VEhBWWTlhI+qXxDFXlejDGcQipMlyxda9nthulaxLZUNcUqFxokp0vcYnvteJln5FNQDRrxj3YcbVw=="
	  },
	  "tslib": {
		"version": "1.10.0",
		"resolved": "https://npm.intra.longguikeji.com/tslib/-/tslib-1.10.0.tgz",
		"integrity": "sha512-qOebF53frne81cf0S9B41ByenJ3/IuH8yJKngAX35CmiZySA0khhkovshKK+jGCaMnVomla7gVlIcc3EvKPbTQ==",
		"dev": true
	  },
	  "tty-browserify": {
		"version": "0.0.0",
		"resolved": "https://npm.intra.longguikeji.com/tty-browserify/-/tty-browserify-0.0.0.tgz",
		"integrity": "sha1-oVe6QC2iTpv5V/mqadUk7tQpAaY=",
		"dev": true
	  },
	  "type-is": {
		"version": "1.6.18",
		"resolved": "https://npm.intra.longguikeji.com/type-is/-/type-is-1.6.18.tgz",
		"integrity": "sha512-TkRKr9sUTxEH8MdfuCSP7VizJyzRNMjj2J2do2Jr3Kym598JVdEksuzPQCnlFPW4ky9Q+iA+ma9BGm06XQBy8g==",
		"requires": {
		  "media-typer": "0.3.0",
		  "mime-types": "~2.1.24"
		}
	  },
	  "typedarray": {
		"version": "0.0.6",
		"resolved": "https://npm.intra.longguikeji.com/typedarray/-/typedarray-0.0.6.tgz",
		"integrity": "sha1-hnrHTjhkGHsdPUfZlqeOxciDB3c=",
		"dev": true
	  },
	  "unicode-canonical-property-names-ecmascript": {
		"version": "1.0.4",
		"resolved": "https://npm.intra.longguikeji.com/unicode-canonical-property-names-ecmascript/-/unicode-canonical-property-names-ecmascript-1.0.4.tgz",
		"integrity": "sha512-jDrNnXWHd4oHiTZnx/ZG7gtUTVp+gCcTTKr8L0HjlwphROEW3+Him+IpvC+xcJEFegapiMZyZe02CyuOnRmbnQ==",
		"dev": true
	  },
	  "unicode-match-property-ecmascript": {
		"version": "1.0.4",
		"resolved": "https://npm.intra.longguikeji.com/unicode-match-property-ecmascript/-/unicode-match-property-ecmascript-1.0.4.tgz",
		"integrity": "sha512-L4Qoh15vTfntsn4P1zqnHulG0LdXgjSO035fEpdtp6YxXhMT51Q6vgM5lYdG/5X3MjS+k/Y9Xw4SFCY9IkR0rg==",
		"dev": true,
		"requires": {
		  "unicode-canonical-property-names-ecmascript": "^1.0.4",
		  "unicode-property-aliases-ecmascript": "^1.0.4"
		}
	  },
	  "unicode-match-property-value-ecmascript": {
		"version": "1.1.0",
		"resolved": "https://npm.intra.longguikeji.com/unicode-match-property-value-ecmascript/-/unicode-match-property-value-ecmascript-1.1.0.tgz",
		"integrity": "sha512-hDTHvaBk3RmFzvSl0UVrUmC3PuW9wKVnpoUDYH0JDkSIovzw+J5viQmeYHxVSBptubnr7PbH2e0fnpDRQnQl5g==",
		"dev": true
	  },
	  "unicode-property-aliases-ecmascript": {
		"version": "1.0.5",
		"resolved": "https://npm.intra.longguikeji.com/unicode-property-aliases-ecmascript/-/unicode-property-aliases-ecmascript-1.0.5.tgz",
		"integrity": "sha512-L5RAqCfXqAwR3RriF8pM0lU0w4Ryf/GgzONwi6KnL1taJQa7x1TCxdJnILX59WIGOwR57IVxn7Nej0fz1Ny6fw==",
		"dev": true
	  },
	  "union-value": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/union-value/-/union-value-1.0.1.tgz",
		"integrity": "sha512-tJfXmxMeWYnczCVs7XAEvIV7ieppALdyepWMkHkwciRpZraG/xwT+s2JN8+pr1+8jCRf80FFzvr+MpQeeoF4Xg==",
		"dev": true,
		"requires": {
		  "arr-union": "^3.1.0",
		  "get-value": "^2.0.6",
		  "is-extendable": "^0.1.1",
		  "set-value": "^2.0.1"
		}
	  },
	  "unique-filename": {
		"version": "1.1.1",
		"resolved": "https://npm.intra.longguikeji.com/unique-filename/-/unique-filename-1.1.1.tgz",
		"integrity": "sha512-Vmp0jIp2ln35UTXuryvjzkjGdRyf9b2lTXuSYUiPmzRcl3FDtYqAwOnTJkAngD9SWhnoJzDbTKwaOrZ+STtxNQ==",
		"dev": true,
		"requires": {
		  "unique-slug": "^2.0.0"
		}
	  },
	  "unique-slug": {
		"version": "2.0.2",
		"resolved": "https://npm.intra.longguikeji.com/unique-slug/-/unique-slug-2.0.2.tgz",
		"integrity": "sha512-zoWr9ObaxALD3DOPfjPSqxt4fnZiWblxHIgeWqW8x7UqDzEtHEQLzji2cuJYQFCU6KmoJikOYAZlrTHHebjx2w==",
		"dev": true,
		"requires": {
		  "imurmurhash": "^0.1.4"
		}
	  },
	  "unpipe": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/unpipe/-/unpipe-1.0.0.tgz",
		"integrity": "sha1-sr9O6FFKrmFltIF4KdIbLvSZBOw="
	  },
	  "unset-value": {
		"version": "1.0.0",
		"resolved": "https://npm.intra.longguikeji.com/unset-value/-/unset-value-1.0.0.tgz",
		"integrity": "sha1-g3aHP30jNRef+x5vw6jtDfyKtVk=",
		"dev": true,
		"requires": {
		  "has-value": "^0.3.1",
		  "isobject": "^3.0.0"
		},
		"dependencies": {
		  "has-value": {
			"version": "0.3.1",
			"resolved": "https://npm.intra.longguikeji.com/has-value/-/has-value-0.3.1.tgz",
			"integrity": "sha1-ex9YutpiyoJ+wKIHgCVlSEWZXh8=",
			"dev": true,
			"requires": {
			  "get-value": "^2.0.3",
			  "has-values": "^0.1.4",
			  "isobject": "^2.0.0"
			},
			"dependencies": {
			  "isobject": {
				"version": "2.1.0",
				"resolved": "https://npm.intra.longguikeji.com/isobject/-/isobject-2.1.0.tgz",
				"integrity": "sha1-8GVWEJaj8dou9GJy+BXIQNh+DIk=",
				"dev": true,
				"requires": {
				  "isarray": "1.0.0"
				}
			  }
			}
		  },
		  "has-values": {
			"version": "0.1.4",
			"resolved": "https://npm.intra.longguikeji.com/has-values/-/has-values-0.1.4.tgz",
			"integrity": "sha1-bWHeldkd/Km5oCCJrThL/49it3E=",
			"dev": true
		  }
		}
	  },
	  "upath": {
		"version": "1.2.0",
		"resolved": "https://npm.intra.longguikeji.com/upath/-/upath-1.2.0.tgz",
		"integrity": "sha512-aZwGpamFO61g3OlfT7OQCHqhGnW43ieH9WZeP7QxN/G/jS4jfqUkZxoryvJgVPEcrl5NL/ggHsSmLMHuH64Lhg==",
		"dev": true
	  },
	  "uri-js": {
		"version": "4.2.2",
		"resolved": "https://npm.intra.longguikeji.com/uri-js/-/uri-js-4.2.2.tgz",
		"integrity": "sha512-KY9Frmirql91X2Qgjry0Wd4Y+YTdrdZheS8TFwvkbLWf/G5KNJDCh6pKL5OZctEW4+0Baa5idK2ZQuELRwPznQ==",
		"dev": true,
		"requires": {
		  "punycode": "^2.1.0"
		}
	  },
	  "urix": {
		"version": "0.1.0",
		"resolved": "https://npm.intra.longguikeji.com/urix/-/urix-0.1.0.tgz",
		"integrity": "sha1-2pN/emLiH+wf0Y1Js1wpNQZ6bHI=",
		"dev": true
	  },
	  "url": {
		"version": "0.11.0",
		"resolved": "https://npm.intra.longguikeji.com/url/-/url-0.11.0.tgz",
		"integrity": "sha1-ODjpfPxgUh63PFJajlW/3Z4uKPE=",
		"dev": true,
		"requires": {
		  "punycode": "1.3.2",
		  "querystring": "0.2.0"
		},
		"dependencies": {
		  "punycode": {
			"version": "1.3.2",
			"resolved": "https://npm.intra.longguikeji.com/punycode/-/punycode-1.3.2.tgz",
			"integrity": "sha1-llOgNvt8HuQjQvIyXM7v6jkmxI0=",
			"dev": true
		  }
		}
	  },
	  "use": {
		"version": "3.1.1",
		"resolved": "https://npm.intra.longguikeji.com/use/-/use-3.1.1.tgz",
		"integrity": "sha512-cwESVXlO3url9YWlFW/TA9cshCEhtu7IKJ/p5soJ/gGpj7vbvFrAY/eIioQ6Dw23KjZhYgiIo8HOs1nQ2vr/oQ==",
		"dev": true
	  },
	  "util": {
		"version": "0.11.1",
		"resolved": "https://npm.intra.longguikeji.com/util/-/util-0.11.1.tgz",
		"integrity": "sha512-HShAsny+zS2TZfaXxD9tYj4HQGlBezXZMZuM/S5PKLLoZkShZiGk9o5CzukI1LVHZvjdvZ2Sj1aW/Ndn2NB/HQ==",
		"dev": true,
		"requires": {
		  "inherits": "2.0.3"
		}
	  },
	  "util-deprecate": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/util-deprecate/-/util-deprecate-1.0.2.tgz",
		"integrity": "sha1-RQ1Nyfpw3nMnYvvS1KKJgUGaDM8=",
		"dev": true
	  },
	  "utils-merge": {
		"version": "1.0.1",
		"resolved": "https://npm.intra.longguikeji.com/utils-merge/-/utils-merge-1.0.1.tgz",
		"integrity": "sha1-n5VxD1CiZ5R7LMwSR0HBAoQn5xM="
	  },
	  "v8-compile-cache": {
		"version": "2.0.3",
		"resolved": "https://npm.intra.longguikeji.com/v8-compile-cache/-/v8-compile-cache-2.0.3.tgz",
		"integrity": "sha512-CNmdbwQMBjwr9Gsmohvm0pbL954tJrNzf6gWL3K+QMQf00PF7ERGrEiLgjuU3mKreLC2MeGhUsNV9ybTbLgd3w==",
		"dev": true
	  },
	  "vary": {
		"version": "1.1.2",
		"resolved": "https://npm.intra.longguikeji.com/vary/-/vary-1.1.2.tgz",
		"integrity": "sha1-IpnwLG3tMNSllhsLn3RSShj2NPw="
	  },
	  "vm-browserify": {
		"version": "1.1.2",
		"resolved": "https://npm.intra.longguikeji.com/vm-browserify/-/vm-browserify-1.1.2.tgz",
		"integrity": "sha512-2ham8XPWTONajOR0ohOKOHXkm3+gaBmGut3SRuu75xLd/RRaY6vqgh8NBYYk7+RW3u5AtzPQZG8F10LHkl0lAQ==",
		"dev": true
	  },
	  "watchpack": {
		"version": "1.6.0",
		"resolved": "https://npm.intra.longguikeji.com/watchpack/-/watchpack-1.6.0.tgz",
		"integrity": "sha512-i6dHe3EyLjMmDlU1/bGQpEw25XSjkJULPuAVKCbNRefQVq48yXKUpwg538F7AZTf9kyr57zj++pQFltUa5H7yA==",
		"dev": true,
		"requires": {
		  "chokidar": "^2.0.2",
		  "graceful-fs": "^4.1.2",
		  "neo-async": "^2.5.0"
		}
	  },
	  "webpack": {
		"version": "4.41.5",
		"resolved": "https://npm.intra.longguikeji.com/webpack/-/webpack-4.41.5.tgz",
		"integrity": "sha512-wp0Co4vpyumnp3KlkmpM5LWuzvZYayDwM2n17EHFr4qxBBbRokC7DJawPJC7TfSFZ9HZ6GsdH40EBj4UV0nmpw==",
		"dev": true,
		"requires": {
		  "@webassemblyjs/ast": "1.8.5",
		  "@webassemblyjs/helper-module-context": "1.8.5",
		  "@webassemblyjs/wasm-edit": "1.8.5",
		  "@webassemblyjs/wasm-parser": "1.8.5",
		  "acorn": "^6.2.1",
		  "ajv": "^6.10.2",
		  "ajv-keywords": "^3.4.1",
		  "chrome-trace-event": "^1.0.2",
		  "enhanced-resolve": "^4.1.0",
		  "eslint-scope": "^4.0.3",
		  "json-parse-better-errors": "^1.0.2",
		  "loader-runner": "^2.4.0",
		  "loader-utils": "^1.2.3",
		  "memory-fs": "^0.4.1",
		  "micromatch": "^3.1.10",
		  "mkdirp": "^0.5.1",
		  "neo-async": "^2.6.1",
		  "node-libs-browser": "^2.2.1",
		  "schema-utils": "^1.0.0",
		  "tapable": "^1.1.3",
		  "terser-webpack-plugin": "^1.4.3",
		  "watchpack": "^1.6.0",
		  "webpack-sources": "^1.4.1"
		}
	  },
	  "webpack-cli": {
		"version": "3.3.10",
		"resolved": "https://npm.intra.longguikeji.com/webpack-cli/-/webpack-cli-3.3.10.tgz",
		"integrity": "sha512-u1dgND9+MXaEt74sJR4PR7qkPxXUSQ0RXYq8x1L6Jg1MYVEmGPrH6Ah6C4arD4r0J1P5HKjRqpab36k0eIzPqg==",
		"dev": true,
		"requires": {
		  "chalk": "2.4.2",
		  "cross-spawn": "6.0.5",
		  "enhanced-resolve": "4.1.0",
		  "findup-sync": "3.0.0",
		  "global-modules": "2.0.0",
		  "import-local": "2.0.0",
		  "interpret": "1.2.0",
		  "loader-utils": "1.2.3",
		  "supports-color": "6.1.0",
		  "v8-compile-cache": "2.0.3",
		  "yargs": "13.2.4"
		},
		"dependencies": {
		  "ansi-regex": {
			"version": "4.1.0",
			"resolved": "https://npm.intra.longguikeji.com/ansi-regex/-/ansi-regex-4.1.0.tgz",
			"integrity": "sha512-1apePfXM1UOSqw0o9IiFAovVz9M5S1Dg+4TrDwfMewQ6p/rmMueb7tWZjQ1rx4Loy1ArBggoqGpfqqdI4rondg==",
			"dev": true
		  },
		  "ansi-styles": {
			"version": "3.2.1",
			"resolved": "https://npm.intra.longguikeji.com/ansi-styles/-/ansi-styles-3.2.1.tgz",
			"integrity": "sha512-VT0ZI6kZRdTh8YyJw3SMbYm/u+NqfsAxEpWO0Pf9sq8/e94WxxOpPKx9FR1FlyCtOVDNOQ+8ntlqFxiRc+r5qA==",
			"dev": true,
			"requires": {
			  "color-convert": "^1.9.0"
			}
		  },
		  "cliui": {
			"version": "5.0.0",
			"resolved": "https://npm.intra.longguikeji.com/cliui/-/cliui-5.0.0.tgz",
			"integrity": "sha512-PYeGSEmmHM6zvoef2w8TPzlrnNpXIjTipYK780YswmIP9vjxmd6Y2a3CB2Ks6/AU8NHjZugXvo8w3oWM2qnwXA==",
			"dev": true,
			"requires": {
			  "string-width": "^3.1.0",
			  "strip-ansi": "^5.2.0",
			  "wrap-ansi": "^5.1.0"
			}
		  },
		  "color-convert": {
			"version": "1.9.3",
			"resolved": "https://npm.intra.longguikeji.com/color-convert/-/color-convert-1.9.3.tgz",
			"integrity": "sha512-QfAUtd+vFdAtFQcC8CCyYt1fYWxSqAiK2cSD6zDB8N3cpsEBAvRxp9zOGg6G/SHHJYAT88/az/IuDGALsNVbGg==",
			"dev": true,
			"requires": {
			  "color-name": "1.1.3"
			}
		  },
		  "color-name": {
			"version": "1.1.3",
			"resolved": "https://npm.intra.longguikeji.com/color-name/-/color-name-1.1.3.tgz",
			"integrity": "sha1-p9BVi9icQveV3UIyj3QIMcpTvCU=",
			"dev": true
		  },
		  "emoji-regex": {
			"version": "7.0.3",
			"resolved": "https://npm.intra.longguikeji.com/emoji-regex/-/emoji-regex-7.0.3.tgz",
			"integrity": "sha512-CwBLREIQ7LvYFB0WyRvwhq5N5qPhc6PMjD6bYggFlI5YyDgl+0vxq5VHbMOFqLg7hfWzmu8T5Z1QofhmTIhItA==",
			"dev": true
		  },
		  "enhanced-resolve": {
			"version": "4.1.0",
			"resolved": "https://npm.intra.longguikeji.com/enhanced-resolve/-/enhanced-resolve-4.1.0.tgz",
			"integrity": "sha512-F/7vkyTtyc/llOIn8oWclcB25KdRaiPBpZYDgJHgh/UHtpgT2p2eldQgtQnLtUvfMKPKxbRaQM/hHkvLHt1Vng==",
			"dev": true,
			"requires": {
			  "graceful-fs": "^4.1.2",
			  "memory-fs": "^0.4.0",
			  "tapable": "^1.0.0"
			}
		  },
		  "find-up": {
			"version": "3.0.0",
			"resolved": "https://npm.intra.longguikeji.com/find-up/-/find-up-3.0.0.tgz",
			"integrity": "sha512-1yD6RmLI1XBfxugvORwlck6f75tYL+iR0jqwsOrOxMZyGYqUuDhJ0l4AXdO1iX/FTs9cBAMEk1gWSEx1kSbylg==",
			"dev": true,
			"requires": {
			  "locate-path": "^3.0.0"
			}
		  },
		  "is-fullwidth-code-point": {
			"version": "2.0.0",
			"resolved": "https://npm.intra.longguikeji.com/is-fullwidth-code-point/-/is-fullwidth-code-point-2.0.0.tgz",
			"integrity": "sha1-o7MKXE8ZkYMWeqq5O+764937ZU8=",
			"dev": true
		  },
		  "locate-path": {
			"version": "3.0.0",
			"resolved": "https://npm.intra.longguikeji.com/locate-path/-/locate-path-3.0.0.tgz",
			"integrity": "sha512-7AO748wWnIhNqAuaty2ZWHkQHRSNfPVIsPIfwEOWO22AmaoVrWavlOcMR5nzTLNYvp36X220/maaRsrec1G65A==",
			"dev": true,
			"requires": {
			  "p-locate": "^3.0.0",
			  "path-exists": "^3.0.0"
			}
		  },
		  "p-locate": {
			"version": "3.0.0",
			"resolved": "https://npm.intra.longguikeji.com/p-locate/-/p-locate-3.0.0.tgz",
			"integrity": "sha512-x+12w/To+4GFfgJhBEpiDcLozRJGegY+Ei7/z0tSLkMmxGZNybVMSfWj9aJn8Z5Fc7dBUNJOOVgPv2H7IwulSQ==",
			"dev": true,
			"requires": {
			  "p-limit": "^2.0.0"
			}
		  },
		  "path-exists": {
			"version": "3.0.0",
			"resolved": "https://npm.intra.longguikeji.com/path-exists/-/path-exists-3.0.0.tgz",
			"integrity": "sha1-zg6+ql94yxiSXqfYENe1mwEP1RU=",
			"dev": true
		  },
		  "string-width": {
			"version": "3.1.0",
			"resolved": "https://npm.intra.longguikeji.com/string-width/-/string-width-3.1.0.tgz",
			"integrity": "sha512-vafcv6KjVZKSgz06oM/H6GDBrAtz8vdhQakGjFIvNrHA6y3HCF1CInLy+QLq8dTJPQ1b+KDUqDFctkdRW44e1w==",
			"dev": true,
			"requires": {
			  "emoji-regex": "^7.0.1",
			  "is-fullwidth-code-point": "^2.0.0",
			  "strip-ansi": "^5.1.0"
			}
		  },
		  "strip-ansi": {
			"version": "5.2.0",
			"resolved": "https://npm.intra.longguikeji.com/strip-ansi/-/strip-ansi-5.2.0.tgz",
			"integrity": "sha512-DuRs1gKbBqsMKIZlrffwlug8MHkcnpjs5VPmL1PAh+mA30U0DTotfDZ0d2UUsXpPmPmMMJ6W773MaA3J+lbiWA==",
			"dev": true,
			"requires": {
			  "ansi-regex": "^4.1.0"
			}
		  },
		  "supports-color": {
			"version": "6.1.0",
			"resolved": "https://npm.intra.longguikeji.com/supports-color/-/supports-color-6.1.0.tgz",
			"integrity": "sha512-qe1jfm1Mg7Nq/NSh6XE24gPXROEVsWHxC1LIx//XNlD9iw7YZQGjZNjYN7xGaEG6iKdA8EtNFW6R0gjnVXp+wQ==",
			"dev": true,
			"requires": {
			  "has-flag": "^3.0.0"
			}
		  },
		  "wrap-ansi": {
			"version": "5.1.0",
			"resolved": "https://npm.intra.longguikeji.com/wrap-ansi/-/wrap-ansi-5.1.0.tgz",
			"integrity": "sha512-QC1/iN/2/RPVJ5jYK8BGttj5z83LmSKmvbvrXPNCLZSEb32KKVDJDl/MOt2N01qU2H/FkzEa9PKto1BqDjtd7Q==",
			"dev": true,
			"requires": {
			  "ansi-styles": "^3.2.0",
			  "string-width": "^3.0.0",
			  "strip-ansi": "^5.0.0"
			}
		  },
		  "yargs": {
			"version": "13.2.4",
			"resolved": "https://npm.intra.longguikeji.com/yargs/-/yargs-13.2.4.tgz",
			"integrity": "sha512-HG/DWAJa1PAnHT9JAhNa8AbAv3FPaiLzioSjCcmuXXhP8MlpHO5vwls4g4j6n30Z74GVQj8Xa62dWVx1QCGklg==",
			"dev": true,
			"requires": {
			  "cliui": "^5.0.0",
			  "find-up": "^3.0.0",
			  "get-caller-file": "^2.0.1",
			  "os-locale": "^3.1.0",
			  "require-directory": "^2.1.1",
			  "require-main-filename": "^2.0.0",
			  "set-blocking": "^2.0.0",
			  "string-width": "^3.0.0",
			  "which-module": "^2.0.0",
			  "y18n": "^4.0.0",
			  "yargs-parser": "^13.1.0"
			}
		  },
		  "yargs-parser": {
			"version": "13.1.1",
			"resolved": "https://npm.intra.longguikeji.com/yargs-parser/-/yargs-parser-13.1.1.tgz",
			"integrity": "sha512-oVAVsHz6uFrg3XQheFII8ESO2ssAf9luWuAd6Wexsu4F3OtIW0o8IribPXYrD4WC24LWtPrJlGy87y5udK+dxQ==",
			"dev": true,
			"requires": {
			  "camelcase": "^5.0.0",
			  "decamelize": "^1.2.0"
			}
		  }
		}
	  },
	  "webpack-sources": {
		"version": "1.4.3",
		"resolved": "https://npm.intra.longguikeji.com/webpack-sources/-/webpack-sources-1.4.3.tgz",
		"integrity": "sha512-lgTS3Xhv1lCOKo7SA5TjKXMjpSM4sBjNV5+q2bqesbSPs5FjGmU6jjtBSkX9b4qW87vDIsCIlUPOEhbZrMdjeQ==",
		"dev": true,
		"requires": {
		  "source-list-map": "^2.0.0",
		  "source-map": "~0.6.1"
		},
		"dependencies": {
		  "source-map": {
			"version": "0.6.1",
			"resolved": "https://npm.intra.longguikeji.com/source-map/-/source-map-0.6.1.tgz",
			"integrity": "sha512-UjgapumWlbMhkBgzT7Ykc5YXUT46F0iKu8SGXq0bcwP5dz/h0Plj6enJqjz1Zbq2l5WaqYnrVbwWOWMyF3F47g==",
			"dev": true
		  }
		}
	  },
	  "which": {
		"version": "1.3.1",
		"resolved": "https://npm.intra.longguikeji.com/which/-/which-1.3.1.tgz",
		"integrity": "sha512-HxJdYWq1MTIQbJ3nw0cqssHoTNU267KlrDuGZ1WYlxDStUtKUhOaJmh112/TZmHxxUfuJqPXSOm7tDyas0OSIQ==",
		"dev": true,
		"requires": {
		  "isexe": "^2.0.0"
		}
	  },
	  "which-module": {
		"version": "2.0.0",
		"resolved": "https://npm.intra.longguikeji.com/which-module/-/which-module-2.0.0.tgz",
		"integrity": "sha1-2e8H3Od7mQK4o6j6SzHD4/fm6Ho="
	  },
	  "worker-farm": {
		"version": "1.7.0",
		"resolved": "https://npm.intra.longguikeji.com/worker-farm/-/worker-farm-1.7.0.tgz",
		"integrity": "sha512-rvw3QTZc8lAxyVrqcSGVm5yP/IJ2UcB3U0graE3LCFoZ0Yn2x4EoVSqJKdB/T5M+FLcRPjz4TDacRf3OCfNUzw==",
		"dev": true,
		"requires": {
		  "errno": "~0.1.7"
		}
	  },
	  "wrap-ansi": {
		"version": "6.2.0",
		"resolved": "https://npm.intra.longguikeji.com/wrap-ansi/-/wrap-ansi-6.2.0.tgz",
		"integrity": "sha512-r6lPcBGxZXlIcymEu7InxDMhdW0KDxpLgoFLcguasxCaJ/SOIZwINatK9KY/tf+ZrlywOKU0UDj3ATXUBfxJXA==",
		"requires": {
		  "ansi-styles": "^4.0.0",
		  "string-width": "^4.1.0",
		  "strip-ansi": "^6.0.0"
		}
	  },
	  "wrappy": {
		"version": "1.0.2",
		"resolved": "https://npm.intra.longguikeji.com/wrappy/-/wrappy-1.0.2.tgz",
		"integrity": "sha1-tSQ9jz7BqjXxNkYFvA0QNuMKtp8=",
		"dev": true
	  },
	  "xtend": {
		"version": "4.0.2",
		"resolved": "https://npm.intra.longguikeji.com/xtend/-/xtend-4.0.2.tgz",
		"integrity": "sha512-LKYU1iAXJXUgAXn9URjiu+MWhyUXHsvfp7mcuYm9dSUKK0/CjtrUwFAxD82/mCWbtLsGjFIad0wIsod4zrTAEQ==",
		"dev": true
	  },
	  "y18n": {
		"version": "4.0.0",
		"resolved": "https://npm.intra.longguikeji.com/y18n/-/y18n-4.0.0.tgz",
		"integrity": "sha512-r9S/ZyXu/Xu9q1tYlpsLIsa3EeLXXk0VwlxqTcFRfg9EhMW+17kbt9G0NrgCmhGb5vT2hyhJZLfDGx+7+5Uj/w=="
	  },
	  "yallist": {
		"version": "3.1.1",
		"resolved": "https://npm.intra.longguikeji.com/yallist/-/yallist-3.1.1.tgz",
		"integrity": "sha512-a4UGQaWPH59mOXUYnAG2ewncQS4i4F43Tv3JoAM+s2VDAmS9NsK8GpDMLrCHPksFT7h3K6TOoUNn2pb7RoXx4g==",
		"dev": true
	  },
	  "yargs": {
		"version": "15.1.0",
		"resolved": "https://npm.intra.longguikeji.com/yargs/-/yargs-15.1.0.tgz",
		"integrity": "sha512-T39FNN1b6hCW4SOIk1XyTOWxtXdcen0t+XYrysQmChzSipvhBO8Bj0nK1ozAasdk24dNWuMZvr4k24nz+8HHLg==",
		"requires": {
		  "cliui": "^6.0.0",
		  "decamelize": "^1.2.0",
		  "find-up": "^4.1.0",
		  "get-caller-file": "^2.0.1",
		  "require-directory": "^2.1.1",
		  "require-main-filename": "^2.0.0",
		  "set-blocking": "^2.0.0",
		  "string-width": "^4.2.0",
		  "which-module": "^2.0.0",
		  "y18n": "^4.0.0",
		  "yargs-parser": "^16.1.0"
		}
	  },
	  "yargs-parser": {
		"version": "16.1.0",
		"resolved": "https://npm.intra.longguikeji.com/yargs-parser/-/yargs-parser-16.1.0.tgz",
		"integrity": "sha512-H/V41UNZQPkUMIT5h5hiwg4QKIY1RPvoBV4XcjUbRM8Bk2oKqqyZ0DIEbTFZB0XjbtSPG8SAa/0DxCQmiRgzKg==",
		"requires": {
		  "camelcase": "^5.0.0",
		  "decamelize": "^1.2.0"
		}
	  }
	}
  }
`
