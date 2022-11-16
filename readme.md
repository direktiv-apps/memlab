
# memlab 1.0

Run memlab in Direktiv on buster

---
- #### Categories: build, development
- #### Image: gcr.io/direktiv/functions/memlab 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/memlab/issues
- #### URL: https://github.com/direktiv-apps/memlab
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About memlab

This function provides a memlab install on Node.js as a Direktiv function. memlab is a memory leak detector for front-end JS.

*NOTE: memlab requires siginificant disk space. Ensure that the following setting has been applied in the `direktiv-config-functions` configmap in your Kubernetes environment:*

```yaml 
# max ephemeral storage in MB
storage: 2048
```

Node Version Manager is installed to support LTS versions. The following versions are installed in this function:

- 18.10.0

- 16.17.1

NVM (Node Version Manager) can be used as well to install different versions but it is function wide which means changes are visible to all function calls during the function / container lifetime. If the application is returning plain JSON on standard out it will be used as JSON result in Direktiv. If the application prints other strings to standard out the response will be a plain string. If JSON output is required the application can create and write to a file called output.json. If this file exists, this function uses its contents as return value.
Functions can have a context to persist the node_modules directory across different execution cycles. Unlike Direktiv's regular behaviour to have a new working directory for each execution, the context ensures that it runs in the same directory each time. 

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: memlab
  image: gcr.io/direktiv/functions/memlab:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: memlab 
  type: action
  action:
    function: memlab
    input:
      files:
      - name: scenario.js
        data: |
          // initial page load's url
          function url() {
            return "https://www.youtube.com";
          }

          // action where you suspect the memory leak might be happening
          async function action(page) {
            await page.click('[id="video-title-link"]');
          }

          // how to go back to the state before action
          async function back(page) {
            await page.click('[id="logo-icon"]');
          }

          module.exports = { action, back, url };
      commands:
      - command: memlab run --scenario scenario.js
```
   #### Change node version
```yaml
- id: memlab 
  type: action
  action:
    function: memlab
    input:
      files:
      - name: scenario.js
        data: |
          // initial page load's url
          function url() {
            return "https://www.youtube.com";
          }

          // action where you suspect the memory leak might be happening
          async function action(page) {
            await page.click('[id="video-title-link"]');
          }

          // how to go back to the state before action
          async function back(page) {
            await page.click('[id="logo-icon"]');
          }

          module.exports = { action, back, url };
      - name: node.js
        data: |
          const {readFileSync} = require('fs')

          const file = readFileSync('/tmp/memlab/data/out/leaks.txt')
          console.log(file.toString())
      commands:
      - command: memlab reset
      - command: memlab run --scenario scenario.js --skip-screenshot --skip-gc --skip-scroll --skip-extra-ops
      - command: node node.js > out/leaks.txt
```
   #### Using a context
```yaml
- id: memlab 
  type: action
  action:
    function: memlab
    input: 
      context: memlab-app
      files: 
      - name: scenario.js
        data: |
          // initial page load's url
          function url() {
            return "https://www.youtube.com";
          }

          // action where you suspect the memory leak might be happening
          async function action(page) {
            await page.click('[id="video-title-link"]');
          }

          // how to go back to the state before action
          async function back(page) {
            await page.click('[id="logo-icon"]');
          }

          module.exports = { action, back, url };
      commands:
      - command: npm install uuid
      - command: memlab run --scenario scenario.js  
```
   #### Using Direktiv variable as script
```yaml
- id: memlab 
  type: action
  action:
    function: memlab
    files:
    - key: scenario.js
      scope: workflow
    input:
      commands:
      - command: memlab run --scenario scenario.js      
```

   ### Secrets


*No secrets required*







### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": "page-load [30.4MB] (baseline) [s1] \u003e action-on-page [45.1MB] (target) [s2] \u003e revert [47.3MB] (final) [s3]  \n------17 clusters------\n\n--Similar leaks in this run: 3308--\n--Retained size of leaked objects: 1.7MB--\n[Window] (native) @161701 [61.3KB]\n  --7 (element)---\u003e  [HTMLDocument] (native) @161699 [29.6KB]\n  --get _activeElement (property)---\u003e  [get activeElement] (closure) @639147 [64 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @460175 [232.4KB]\n  --ug (variable)---\u003e  [WeakMap] (object) @1214207 [131.1KB]\n  --table (internal)---\u003e  [\u003carray\u003e] (array) @1831547 [131KB]\n  --7959 / part of key (f @1517857) -\u003e value (sg @1611255) pair in WeakMap (table @1831547) (internal)---\u003e  [sg] (object) @1611255 [16 bytes]\n  --node (property)---\u003e  [Detached HTMLElement] (native) @1517857 [2.8KB]\n  --69 (element)---\u003e  [Detached InternalNode] (native) @1650270976 [120 bytes]\n  --3 (element)---\u003e  [Detached ElementIntersectionObserverData] (native) @1652327680 [64 bytes]\n\n--Similar leaks in this run: 1582--\n--Retained size of leaked objects: 435.2KB--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --PolymerFakeBaseClassWithoutHtml (property)---\u003e  [\u003cclosure\u003e] (closure) @338929 [224 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @162793 [2.7MB]\n  --guc (variable)---\u003e  [Detached HTMLTemplateElement] (native) @153729 [6.7KB]\n  --15 (element)---\u003e  [Detached DocumentFragment] (native) @1648467776 [5.9KB]\n  --3 (element)---\u003e  [Detached Text] (native) @1648472416 [96 bytes]\n  --2 (element)---\u003e  [Detached HTMLDivElement] (native) @1648471936 [5.4KB]\n  --2 (element)---\u003e  [Detached Text] (native) @1648474816 [96 bytes]\n  --2 (element)---\u003e  [Detached HTMLDivElement] (native) @1648471616 [1.1KB]\n  --4 (element)---\u003e  [Detached Text] (native) @1648472096 [96 bytes]\n  --2 (element)---\u003e  [Detached HTMLDivElement] (native) @1648471776 [1.5KB]\n  --6 (element)---\u003e  [Detached InternalNode] (native) @1087422880 [384 bytes]\n  --3 (element)---\u003e  [Detached InternalNode] (native) @1087422400 [192 bytes]\n  --1 (element)---\u003e  [Detached InternalNode] (native) @1649883680 [192 bytes]\n  --2 (element)---\u003e  [Detached Attr] (native) @1648707072 [96 bytes]\n\n--Similar leaks in this run: 3--\n--Retained size of leaked objects: 6.8KB--\n[Window] (native) @161701 [61.3KB]\n  --7 (element)---\u003e  [HTMLDocument] (native) @161699 [29.6KB]\n  --29 (element)---\u003e  [HTMLAnchorElement] (native) @156491 [348 bytes]\n  --__dataHost (property)---\u003e  [HTMLElement] (native) @156659 [2.8KB]\n  --parentComponent (property)---\u003e  [HTMLElement] (native) @157729 [4.7KB]\n  --__shady (property)---\u003e  [gd] (object) @747797 [264 bytes]\n  --lastChild (property)---\u003e  [HTMLDivElement] (native) @157089 [640 bytes]\n  --__shady (property)---\u003e  [gd] (object) @747615 [32 bytes]\n  --previousSibling (property)---\u003e  [Detached HTMLDivElement] (native) @157101 [14KB]\n  --6 (element)---\u003e  [Detached HTMLSpanElement] (native) @1650970464 [112 bytes]\n  --2 (element)---\u003e  [Detached HTMLAnchorElement] (native) @1650966144 [6.9KB]\n  --2 (element)---\u003e  [Detached SVGSVGElement] (native) @1650965824 [6.7KB]\n  --9 (element)---\u003e  [Detached SVGGElement] (native) @1089022592 [5.7KB]\n  --5 (element)---\u003e  [Detached SVGGElement] (native) @1086709984 [4KB]\n  --4 (element)---\u003e  [Detached SVGPathElement] (native) @1089023552 [464 bytes]\n  --2 (element)---\u003e  [Detached SVGAnimatedNumber] (native) @1632501472 [80 bytes]\n\n--Similar leaks in this run: 12--\n--Retained size of leaked objects: 2.9KB--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --ytPubsubPubsubInstance (property)---\u003e  [yh] (object) @705479 [3.1KB]\n  --subscriptions_ (property)---\u003e  [Array] (object) @717231 [1.7KB]\n  --3 (element)---\u003e  [WU] (object) @314953 [2.6KB]\n  --app (property)---\u003e  [g.iZ] (object) @213143 [158.3KB]\n  --xp (property)---\u003e  [OU] (object) @323541 [84.1KB]\n  --j (property)---\u003e  [Array] (object) @325261 [72.6KB]\n  --776 (element)---\u003e  [native_bind] (closure) @2132673 [26KB]\n  --bound_this (internal)---\u003e  [bN] (object) @2213673 [25.9KB]\n  --T (property)---\u003e  [TM] (object) @2213811 [1.2KB]\n  --element (property)---\u003e  [Detached HTMLDivElement] (native) @2111451 [396 bytes]\n  --5 (element)---\u003e  [Detached Text] (native) @1649255488 [96 bytes]\n\n--Similar leaks in this run: 9--\n--Retained size of leaked objects: 1.9KB--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --Polymer (property)---\u003e  [\u003cclosure\u003e] (closure) @204317 [9.9KB]\n  --telemetry (property)---\u003e  [Object] (object) @204331 [444 bytes]\n  --registrations (property)---\u003e  [Array] (object) @733433 [5.2KB]\n  --6 (element)---\u003e  [a] (object) @733437 [808 bytes]\n  --constructor (property)---\u003e  [f] (closure) @161515 [912 bytes]\n  --generatedFrom (property)---\u003e  [Object] (object) @738589 [3.3KB]\n  --_template (property)---\u003e  [Detached HTMLTemplateElement] (native) @161509 [1KB]\n  --3 (element)---\u003e  [Detached DocumentFragment] (native) @1650932576 [888 bytes]\n  --2 (element)---\u003e  [Detached Text] (native) @1650932256 [96 bytes]\n  --2 (element)---\u003e  [Detached HTMLStyleElement] (native) @1650931776 [304 bytes]\n\n--Similar leaks in this run: 5--\n--Retained size of leaked objects: 1.8KB--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --6 (shortcut)---\u003e  [Window / https://tpc.googlesyndication.com] (object) @2103279 [51.6KB]\n  --botguard (property)---\u003e  [Object] (object) @2190567 [320 bytes]\n  --FDL_ (property)---\u003e  [\u003cclosure\u003e] (closure) @2190569 [68 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @2173673 [18.5KB]\n  --R (variable)---\u003e  [R] (closure) @2223951 [68 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @2202409 [3.4KB]\n  --this (variable)---\u003e  [V] (object) @2162101 [165.8KB]\n  --N (property)---\u003e  [Array] (object) @2179997 [121.6KB]\n  --288 (element)---\u003e  [Object] (object) @2200763 [240 bytes]\n  --concat (property)---\u003e  [\u003cclosure\u003e] (closure) @2204053 [32 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @2202457 [148 bytes]\n  --M (variable)---\u003e  [Array] (object) @2202459 [100 bytes]\n  --6 (element)---\u003e  [Array] (object) @2203783 [36 bytes]\n  --0 (element)---\u003e  [Detached HTMLImageElement] (native) @2111403 [1.1KB]\n  --4 (element)---\u003e  [Detached InternalNode] (native) @1101397568 [864 bytes]\n  --1 (element)---\u003e  [Detached ShadowRoot] (native) @1110164192 [864 bytes]\n  --6 (element)---\u003e  [Detached HTMLSpanElement] (native) @1101393888 [112 bytes]\n  --1 (element)---\u003e  [Detached HTMLImageElement] (native) @1636561024 [216 bytes]\n\n--Similar leaks in this run: 9--\n--Retained size of leaked objects: 1.3KB--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --PolymerFakeBaseClassWithoutHtml (property)---\u003e  [\u003cclosure\u003e] (closure) @338929 [224 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @162793 [2.7MB]\n  --vqb (variable)---\u003e  [HTMLElement] (native) @161857 [17.5KB]\n  --__templateInfo (property)---\u003e  [Object] (object) @597753 [196 bytes]\n  --nodeList (property)---\u003e  [Array] (object) @850857 [96 bytes]\n  --8 (element)---\u003e  [HTMLTemplateElement] (native) @157341 [284 bytes]\n  --_templateInfo (property)---\u003e  [Object] (object) @364877 [4.4KB]\n  --content (property)---\u003e  [Detached DocumentFragment] (native) @156951 [4KB]\n  --4 (element)---\u003e  [Detached Text] (native) @1639825344 [128 bytes]\n  --2 (element)---\u003e  [Detached HTMLElement] (native) @1639833664 [3.6KB]\n  --6 (element)---\u003e  [Detached InternalNode] (native) @1085199296 [384 bytes]\n  --3 (element)---\u003e  [Detached InternalNode] (native) @1085198816 [192 bytes]\n  --1 (element)---\u003e  [Detached InternalNode] (native) @1649903648 [192 bytes]\n  --1 (element)---\u003e  [Detached Attr] (native) @1649372832 [96 bytes]\n\n--Similar leaks in this run: 4--\n--Retained size of leaked objects: 984 bytes--\n[Window] (native) @161701 [61.3KB]\n  --7 (element)---\u003e  [HTMLDocument] (native) @161699 [29.6KB]\n  --get _activeElement (property)---\u003e  [get activeElement] (closure) @639147 [64 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @460175 [232.4KB]\n  --oi (variable)---\u003e  [Object] (object) @1213899 [3.1KB]\n  --iron-a11y-announcer (property)---\u003e  [Detached HTMLTemplateElement] (native) @161467 [1.6KB]\n  --_styles (property)---\u003e  [Array] (object) @424477 [584 bytes]\n  --0 (element)---\u003e  [Detached HTMLStyleElement] (native) @154797 [492 bytes]\n  --4 (element)---\u003e  [Detached Text] (native) @1638671424 [96 bytes]\n\n--Similar leaks in this run: 3--\n--Retained size of leaked objects: 740 bytes--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --PolymerFakeBaseClassWithoutHtml (property)---\u003e  [\u003cclosure\u003e] (closure) @338929 [224 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @162793 [2.7MB]\n  --jQb (variable)---\u003e  [Set] (object) @1186849 [356 bytes]\n  --table (internal)---\u003e  [\u003carray\u003e] (array) @1610289 [340 bytes]\n  --33 (internal)---\u003e  [HTMLElement] (native) @1528041 [3.7KB]\n  --__shady (property)---\u003e  [gd] (object) @1652243 [228 bytes]\n  --lastChild (property)---\u003e  [Detached Text] (native) @1527769 [176 bytes]\n\n--Similar leaks in this run: 5--\n--Retained size of leaked objects: 512 bytes--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --PolymerFakeBaseClassWithoutHtml (property)---\u003e  [\u003cclosure\u003e] (closure) @338929 [224 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @162793 [2.7MB]\n  --VYc (variable)---\u003e  [UYc] (object) @591515 [14.5KB]\n  --audioPlayer (property)---\u003e  [NYc] (object) @961303 [11.7KB]\n  --audioFeedbackHolder (property)---\u003e  [Object] (object) @961305 [11.5KB]\n  --success (property)---\u003e  [Detached HTMLAudioElement] (native) @157723 [2.8KB]\n  --13 (element)---\u003e  [Detached InternalNode] (native) @1649917248 [328 bytes]\n  --1 (element)---\u003e  [Detached ShadowRoot] (native) @1649818144 [328 bytes]\n\n--Similar leaks in this run: 1--\n--Retained size of leaked objects: 328 bytes--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --yt (property)---\u003e  [Object] (object) @641361 [9.9KB]\n  --player (property)---\u003e  [Object] (object) @1140291 [3.5KB]\n  --utils (property)---\u003e  [Object] (object) @1192601 [3.3KB]\n  --videoElement_ (property)---\u003e  [Detached HTMLVideoElement] (native) @123867 [3.2KB]\n  --11 (element)---\u003e  [Detached InternalNode] (native) @1084113312 [392 bytes]\n  --1 (element)---\u003e  [Detached ShadowRoot] (native) @1084113792 [328 bytes]\n\n--Similar leaks in this run: 3--\n--Retained size of leaked objects: 272 bytes--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --trayride (property)---\u003e  [Object] (object) @353529 [320 bytes]\n  --kDx_ (property)---\u003e  [\u003cclosure\u003e] (closure) @682905 [68 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @912549 [13.6KB]\n  --uZ (variable)---\u003e  [uZ] (closure) @896177 [68 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @896061 [1KB]\n  --this (variable)---\u003e  [I] (object) @896021 [307.7KB]\n  --J (property)---\u003e  [Array] (object) @1648491 [243.5KB]\n  --128 (element)---\u003e  [Object] (object) @1133215 [86.9KB]\n  --concat (property)---\u003e  [\u003cclosure\u003e] (closure) @1134939 [32 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @1134931 [86.9KB]\n  --Y (variable)---\u003e  [Array] (object) @1134933 [86.8KB]\n  --0 (element)---\u003e  [Detached Window] (native) @141183 [66.8KB]\n  --17 (element)---\u003e  [Detached InternalNode] (native) @1088787616 [1.5KB]\n  --3 (element)---\u003e  [Detached InternalNode] (native) @1088792256 [936 bytes]\n  --1 (element)---\u003e  [Detached Performance] (native) @1100982272 [936 bytes]\n  --1 (element)---\u003e  [Detached PerformanceTiming] (native) @1634372576 [48 bytes]\n\n--Similar leaks in this run: 1--\n--Retained size of leaked objects: 244 bytes--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --google_image_requests (property)---\u003e  [Array] (object) @1541401 [336 bytes]\n  --0 (element)---\u003e  [Detached HTMLImageElement] (native) @1523967 [244 bytes]\n\n--Similar leaks in this run: 2--\n--Retained size of leaked objects: 192 bytes--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --ytPubsubPubsubInstance (property)---\u003e  [yh] (object) @705479 [3.1KB]\n  --subscriptions_ (property)---\u003e  [Array] (object) @717231 [1.7KB]\n  --27 (element)---\u003e  [xlb] (object) @318379 [1.9KB]\n  --wl (property)---\u003e  [Array] (object) @1346475 [444 bytes]\n  --1 (element)---\u003e  [\u003cclosure\u003e] (closure) @328923 [88 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @328931 [56 bytes]\n  --b (variable)---\u003e  [m$] (object) @315571 [3.2KB]\n  --B (property)---\u003e  [g.JQ] (object) @315567 [3.1KB]\n  --wc (property)---\u003e  [Object] (object) @1347229 [248 bytes]\n  --ytp-panel-title (property)---\u003e  [Detached HTMLSpanElement] (native) @133227 [324 bytes]\n  --3 (element)---\u003e  [Detached Text] (native) @1649073280 [96 bytes]\n\n--Similar leaks in this run: 2--\n--Retained size of leaked objects: 160 bytes--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --webpocb (property)---\u003e  [J] (closure) @915643 [84 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @915649 [52 bytes]\n  --S (variable)---\u003e  [I] (object) @318817 [166.7KB]\n  --J (property)---\u003e  [Array] (object) @322957 [136.7KB]\n  --116 (element)---\u003e  [Object] (object) @924839 [444 bytes]\n  --concat (property)---\u003e  [\u003cclosure\u003e] (closure) @925281 [32 bytes]\n  --context (internal)---\u003e  [\u003cfunction scope\u003e] (object) @913379 [352 bytes]\n  --Y (variable)---\u003e  [Array] (object) @913381 [304 bytes]\n  --2 (element)---\u003e  [Window] (object) @131373 [240 bytes]\n  --map (internal)---\u003e  [system / Map] (hidden) @1298135 [188.1KB]\n  --prototype (internal)---\u003e  [Window / https://www.youtube.com] (object) @1298137 [187.9KB]\n  --\u003csymbol Window#DocumentCachedAccessor\u003e (property)---\u003e  [Detached HTMLDocument] (native) @131371 [82.7KB]\n  --16 (element)---\u003e  [Detached ScriptedAnimationController] (native) @1101476832 [224 bytes]\n  --1 (element)---\u003e  [Detached Window] (native) @131377 [66.5KB]\n  --17 (element)---\u003e  [Detached InternalNode] (native) @1086469888 [1.5KB]\n  --3 (element)---\u003e  [Detached InternalNode] (native) @1086469568 [936 bytes]\n  --1 (element)---\u003e  [Detached Performance] (native) @1088242080 [936 bytes]\n  --1 (element)---\u003e  [Detached PerformanceTiming] (native) @1634139712 [48 bytes]\n\n--Similar leaks in this run: 1--\n--Retained size of leaked objects: 136 bytes--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --ytPubsubPubsubInstance (property)---\u003e  [yh] (object) @705479 [3.1KB]\n  --subscriptions_ (property)---\u003e  [Array] (object) @717231 [1.7KB]\n  --3 (element)---\u003e  [WU] (object) @314953 [2.6KB]\n  --app (property)---\u003e  [g.iZ] (object) @213143 [158.3KB]\n  --kZ (property)---\u003e  [hJa] (object) @323551 [3.3KB]\n  --I (property)---\u003e  [oP] (object) @325347 [756 bytes]\n  --j (property)---\u003e  [nP] (object) @315435 [1.8KB]\n  --B (property)---\u003e  [g.GT] (object) @315395 [11KB]\n  --Pf (property)---\u003e  [wT] (object) @330365 [8.6KB]\n  --Wa (property)---\u003e  [g.HR] (object) @1346951 [2.2KB]\n  --ctx (property)---\u003e  [Detached CanvasRenderingContext2D] (native) @124903 [900 bytes]\n  --5 (element)---\u003e  [Detached InternalNode] (native) @1084180608 [136 bytes]\n  --1 (element)---\u003e  [Detached InternalNode] (native) @1102152320 [136 bytes]\n  --1 (element)---\u003e  [Detached InternalNode] (native) @1102152160 [136 bytes]\n  --1 (element)---\u003e  [Detached CanvasGradient] (native) @1090642464 [136 bytes]\n\n--Similar leaks in this run: 1--\n--Retained size of leaked objects: 72 bytes--\n[\u003csynthetic\u003e] (synthetic) @1 [55.3MB]\n  --2 (shortcut)---\u003e  [Window / https://www.youtube.com] (object) @9819 [81.2KB]\n  --ytEventsEventsListeners (property)---\u003e  [Object] (object) @312859 [1.6KB]\n  --18 (element)---\u003e  [Array] (object) @1227801 [72 bytes]\n  --0 (element)---\u003e  [HTMLInputElement] (native) @157201 [972 bytes]\n  --8 (element)---\u003e  [HTMLDivElement] (native) @157091 [2.1KB]\n  --__shady (property)---\u003e  [gd] (object) @747619 [1.9KB]\n  --nextSibling (property)---\u003e  [Detached SVGSVGElement] (native) @157099 [1.9KB]\n  --12 (element)---\u003e  [Detached SVGAnimatedString] (native) @1636159488 [72 bytes]",
    "success": true
  }
]
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| memlab | [][PostOKBodyMemlabItems](#post-o-k-body-memlab-items)| `[]*PostOKBodyMemlabItems` |  | |  |  |


#### <span id="post-o-k-body-memlab-items"></span> postOKBodyMemlabItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | `[{"command":"memlab run --scenario scenario.js"}]`| Array of commands. |  |
| context | string| `string` |  | | Direktiv will delete the working directory after each execution. With the context the application can run in a different
directory and commands like npm install will be persistent. If context is not set the "node_module" directory will be deleted
and each execution of the flow uses an empty modules folder. Multiple apps can not share a context. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |
| node | string| `string` |  | `"18.10.0"`| Default node version for the script |  |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run |  |
| continue | boolean| `bool` |  | | Stops excecution if command fails, otherwise proceeds with next command |  |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |

 
