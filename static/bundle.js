!function(t){var e={};function n(o){if(e[o])return e[o].exports;var l=e[o]={i:o,l:!1,exports:{}};return t[o].call(l.exports,l,l.exports,n),l.l=!0,l.exports}n.m=t,n.c=e,n.d=function(t,e,o){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:o})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var o=Object.create(null);if(n.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var l in t)n.d(o,l,function(e){return t[e]}.bind(null,l));return o},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="",n(n.s=0)}([function(t,e,n){t.exports=n(2)},function(t,e,n){},function(t,e,n){"use strict";function o(){}n.r(e);function l(t){return t()}function i(){return Object.create(null)}function r(t){t.forEach(l)}function c(t){return"function"==typeof t}function s(t,e){return t!=t?e==e:t!==e||t&&"object"==typeof t||"function"==typeof t}new Set;function a(t,e){t.appendChild(e)}function u(t,e,n){t.insertBefore(e,n||null)}function f(t){t.parentNode.removeChild(t)}function d(t){return document.createElement(t)}function p(t){return document.createTextNode(t)}function h(){return p(" ")}function g(t,e,n,o){return t.addEventListener(e,n,o),()=>t.removeEventListener(e,n,o)}function m(t,e,n){null==n?t.removeAttribute(e):t.getAttribute(e)!==n&&t.setAttribute(e,n)}function v(t,e){e=""+e,t.data!==e&&(t.data=e)}function b(t,e){(null!=e||t.value)&&(t.value=e)}let y;function $(t){y=t}const k=[],_=[],w=[],x=[],L=Promise.resolve();let T=!1;function j(){T||(T=!0,L.then(M))}function C(t){w.push(t)}let S=!1;const E=new Set;function M(){if(!S){S=!0;do{for(let t=0;t<k.length;t+=1){const e=k[t];$(e),O(e.$$)}for(k.length=0;_.length;)_.pop()();for(let t=0;t<w.length;t+=1){const e=w[t];E.has(e)||(E.add(e),e())}w.length=0}while(k.length);for(;x.length;)x.pop()();T=!1,S=!1,E.clear()}}function O(t){if(null!==t.fragment){t.update(),r(t.before_update);const e=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,e),t.after_update.forEach(C)}}const P=new Set;function q(t,e){t&&t.i&&(P.delete(t),t.i(e))}"undefined"!=typeof window?window:global;new Set(["allowfullscreen","allowpaymentrequest","async","autofocus","autoplay","checked","controls","default","defer","disabled","formnovalidate","hidden","ismap","loop","multiple","muted","nomodule","novalidate","open","playsinline","readonly","required","reversed","selected"]);let A;function H(t,e){const n=t.$$;null!==n.fragment&&(r(n.on_destroy),n.fragment&&n.fragment.d(e),n.on_destroy=n.fragment=null,n.ctx=[])}function G(t,e,n,s,a,u,f=[-1]){const d=y;$(t);const p=e.props||{},h=t.$$={fragment:null,ctx:null,props:u,update:o,not_equal:a,bound:i(),on_mount:[],on_destroy:[],before_update:[],after_update:[],context:new Map(d?d.$$.context:[]),callbacks:i(),dirty:f};let g=!1;var m;h.ctx=n?n(t,p,(e,n,...o)=>{const l=o.length?o[0]:n;return h.ctx&&a(h.ctx[e],h.ctx[e]=l)&&(h.bound[e]&&h.bound[e](l),g&&function(t,e){-1===t.$$.dirty[0]&&(k.push(t),j(),t.$$.dirty.fill(0)),t.$$.dirty[e/31|0]|=1<<e%31}(t,e)),n}):[],h.update(),g=!0,r(h.before_update),h.fragment=!!s&&s(h.ctx),e.target&&(e.hydrate?h.fragment&&h.fragment.l((m=e.target,Array.from(m.childNodes))):h.fragment&&h.fragment.c(),e.intro&&q(t.$$.fragment),function(t,e,n){const{fragment:o,on_mount:i,on_destroy:s,after_update:a}=t.$$;o&&o.m(e,n),C(()=>{const e=i.map(l).filter(c);s?s.push(...e):r(e),t.$$.on_mount=[]}),a.forEach(C)}(t,e.target,e.anchor),M()),$(d)}"function"==typeof HTMLElement&&(A=class extends HTMLElement{constructor(){super(),this.attachShadow({mode:"open"})}connectedCallback(){for(const t in this.$$.slotted)this.appendChild(this.$$.slotted[t])}attributeChangedCallback(t,e,n){this[t]=n}$destroy(){H(this,1),this.$destroy=o}$on(t,e){const n=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return n.push(e),()=>{const t=n.indexOf(e);-1!==t&&n.splice(t,1)}}$set(){}});class N{$destroy(){H(this,1),this.$destroy=o}$on(t,e){const n=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return n.push(e),()=>{const t=n.indexOf(e);-1!==t&&n.splice(t,1)}}$set(){}}n(1);function B(t,e,n){const o=t.slice();return o[5]=e[n],o[15]=n,o}function D(t){let e,n;return{c(){e=d("div"),n=p(t[2]),m(e,"class","flash-message svelte-1guvik7")},m(t,o){u(t,e,o),a(e,n)},p(t,e){4&e&&v(n,t[2])},d(t){t&&f(e)}}}function I(t){let e,n,o,l,i,c,s=t[5]+"";return{c(){e=d("button"),n=p(s),o=h(),m(e,"id",l=`possibleLocation--${t[15]}`),m(e,"name",i=t[5]),m(e,"class","svelte-1guvik7")},m(l,i){u(l,e,i),a(e,n),a(e,o),c=[g(e,"click",t[7]),g(e,"keyup",t[8])]},p(t,o){8&o&&s!==(s=t[5]+"")&&v(n,s),8&o&&i!==(i=t[5])&&m(e,"name",i)},d(t){t&&f(e),r(c)}}}function F(t){let e,n;return{c(){e=d("a"),n=p("Link here!"),m(e,"class","link svelte-1guvik7"),m(e,"href",t[1]),m(e,"target","_blank")},m(t,o){u(t,e,o),a(e,n)},p(t,n){2&n&&m(e,"href",t[1])},d(t){t&&f(e)}}}function K(t){let e,n,l,i,c,s,p,v,y,$,k,_,w,x,L,T,j,C,S,E,M,O,P,q,A,H,G,N,K,R=t[2]&&D(t),W=t[3],z=[];for(let e=0;e<W.length;e+=1)z[e]=I(B(t,W,e));let J=t[1]&&F(t);return{c(){R&&R.c(),e=h(),n=d("div"),l=d("div"),l.innerHTML='<h2 class="svelte-1guvik7">Google Location Faker</h2> \n\t\t<p>See local map results as if you were in the location.</p>',i=h(),c=d("div"),s=d("label"),p=d("input"),v=h(),y=d("span"),y.textContent="Keyword",$=h(),k=d("div"),_=d("label"),w=d("input"),x=h(),L=d("span"),L.textContent="Location",T=h(),j=d("div");for(let t=0;t<z.length;t+=1)z[t].c();S=h(),E=d("button"),E.textContent="Get Results",M=h(),O=d("div"),J&&J.c(),P=h(),q=d("div"),A=d("div"),A.innerHTML='<h2 style="margin-top: 0">How does it work?</h2>\n\t\t\tGoogle uses an algorithm, which is based on the length of the input location, to encode the location data of a search. <br><br> We can fake the output of this algorithm to trick Google to show us local listings. \n\t\t',H=h(),G=d("div"),G.innerHTML="\n\t\t\tThe location input features an optional location autocomplete. This is useful if you are searching for somewhere that has a duplicate name.<br><br> The list of autocomplete names is provided via Google &amp; is case sensitive.\n\t\t",m(l,"class","header svelte-1guvik7"),m(p,"autocomplete","off"),m(p,"placeholder","Keyword"),p.required=!0,m(p,"class","svelte-1guvik7"),m(y,"class","floating-label__el"),m(s,"class","floating-label svelte-1guvik7"),m(w,"placeholder","Location"),w.required=!0,m(w,"class","svelte-1guvik7"),m(L,"class","floating-label__el"),m(_,"class","floating-label"),m(j,"class",C="possible-location-container "+(t[4]?null:"possible-location-container--hide")+" svelte-1guvik7"),m(k,"class","location location-input__container svelte-1guvik7"),m(E,"class","submit svelte-1guvik7"),m(c,"class","input-container svelte-1guvik7"),m(O,"class","link-container svelte-1guvik7"),m(A,"class","instructions__left svelte-1guvik7"),m(G,"class","instructions__right svelte-1guvik7"),m(q,"class",N="instructions "+(t[4]?"instructions--hide":null)+"  svelte-1guvik7"),m(n,"class","container svelte-1guvik7")},m(o,r){R&&R.m(o,r),u(o,e,r),u(o,n,r),a(n,l),a(n,i),a(n,c),a(c,s),a(s,p),b(p,t[0]),a(s,v),a(s,y),a(c,$),a(c,k),a(k,_),a(_,w),b(w,t[5]),a(_,x),a(_,L),a(k,T),a(k,j);for(let t=0;t<z.length;t+=1)z[t].m(j,null);a(c,S),a(c,E),a(n,M),a(n,O),J&&J.m(O,null),a(n,P),a(n,q),a(q,A),a(q,H),a(q,G),K=[g(p,"input",t[12]),g(w,"input",t[13]),g(w,"keyup",t[8]),g(E,"click",t[6])]},p(t,[n]){if(t[2]?R?R.p(t,n):(R=D(t),R.c(),R.m(e.parentNode,e)):R&&(R.d(1),R=null),1&n&&p.value!==t[0]&&b(p,t[0]),32&n&&w.value!==t[5]&&b(w,t[5]),392&n){let e;for(W=t[3],e=0;e<W.length;e+=1){const o=B(t,W,e);z[e]?z[e].p(o,n):(z[e]=I(o),z[e].c(),z[e].m(j,null))}for(;e<z.length;e+=1)z[e].d(1);z.length=W.length}16&n&&C!==(C="possible-location-container "+(t[4]?null:"possible-location-container--hide")+" svelte-1guvik7")&&m(j,"class",C),t[1]?J?J.p(t,n):(J=F(t),J.c(),J.m(O,null)):J&&(J.d(1),J=null),16&n&&N!==(N="instructions "+(t[4]?"instructions--hide":null)+"  svelte-1guvik7")&&m(q,"class",N)},i:o,o:o,d(t){R&&R.d(t),t&&f(e),t&&f(n),function(t,e){for(let n=0;n<t.length;n+=1)t[n]&&t[n].d(e)}(z,t),J&&J.d(),r(K)}}}function R(t,e,n){let o,l,i,r,c="Fr",s=[],a=-1,u=!1,f=!1;return[l,i,r,s,f,o,function(){n(4,f=!1),o&&l&&o.length>3?fetch(`https://location-faker.onrender.com/api?q=${l}&location=${o}`).then(t=>t.json()).then(t=>{n(1,i=`https://google.com/search?${t}`),n(2,r=null)}).catch(t=>{n(2,r="An issue occurred while communicating with the API.")}):o.length<4?n(2,r="Location must be more than 3 characters"):n(2,r="Please fill in the required fields")},function(t){t.preventDefault(),n(5,o=t.currentTarget.name),n(4,f=!1)},function(t){n(1,i=""),n(4,f=!0),38==t.keyCode?-1!==a&&(t.preventDefault(),a-=1,document.getElementById(`possibleLocation--${a}`).focus()):40==t.keyCode?(t.preventDefault(),a+=1,document.getElementById(`possibleLocation--${a}`).focus()):(n(3,s.length=0,s),a=-1,fetch(`https://location-faker.onrender.com/autocomplete?input=${o}`).then(t=>(u=!1,t.json())).then(t=>{c="",n(3,s.length=0,s);for(let e=0;e<4;e+=1)t[e]&&s.push(`${o}${t[e]}`)}).catch(t=>{n(4,f=!1)}))},c,a,u,function(){l=this.value,n(0,l)},function(){o=this.value,n(5,o)}]}const W=new class extends N{constructor(t){super(),G(this,t,R,K,s,{})}}({target:document.body,props:{name:"world"}});window.app=W;e.default=W}]);
