import{_ as D}from"./index.f0008794.js";import{D as B,F as E,I as g,w as A}from"./element-plus.078d4249.js";import{_ as b}from"./picker.26c88e51.js";import{a as h,b as k}from"./user.6e5e2c24.js";import{f as w}from"./index.1eff73d9.js";import{d as n,a1 as y,an as V,o as p,c as x,X as e,P as r,u as a,a as u,Q as S,O as U,U as N}from"./@vue.a137a740.js";import"./@vueuse.07613b64.js";import"./@element-plus.3660753f.js";import"./lodash-es.a31ceab4.js";import"./dayjs.bd523028.js";import"./axios.d8168cfd.js";import"./async-validator.fb49d0f5.js";import"./@ctrl.fd318bfa.js";import"./@popperjs.36402333.js";import"./escape-html.e5dfadb9.js";import"./normalize-wheel-es.8aeb3683.js";import"./index.0b0483de.js";import"./index.7f4255c7.js";import"./usePaging.dfd0c037.js";import"./index.ee2b7eb3.js";import"./index.vue_vue_type_script_setup_true_lang.6e80abd6.js";import"./vue3-video-play.b1eef99b.js";import"./vuedraggable.0ab4ab66.js";import"./vue.e5a65d9e.js";import"./sortablejs.fd7ddf86.js";import"./lodash.48927ea5.js";import"./vue-router.9605b890.js";import"./pinia.9b4180ce.js";import"./css-color-function.b4c88e1a.js";import"./color.a9016252.js";import"./clone.73d1916b.js";import"./color-convert.755d189f.js";import"./color-name.e7a4e1d3.js";import"./color-string.e356f5de.js";import"./balanced-match.d2a36341.js";import"./ms.564e106c.js";import"./nprogress.0f0f7ca7.js";import"./vue-clipboard3.4e164ffd.js";import"./clipboard.7c3d630c.js";import"./echarts.7e912674.js";import"./zrender.754e8e90.js";import"./tslib.60310f1a.js";import"./highlight.js.7165574c.js";import"./@highlightjs.7fc78ec7.js";const j={class:"user-setup"},I=u("div",{class:"font-medium mb-7"},"\u57FA\u672C\u8BBE\u7F6E",-1),O=u("div",null,[u("div",{class:"form-tips"}," \u7528\u6237\u6CE8\u518C\u65F6\u7ED9\u7684\u9ED8\u8BA4\u5934\u50CF\uFF0C\u5EFA\u8BAE\u5C3A\u5BF8\uFF1A400*400\u50CF\u7D20\uFF0C\u652F\u6301jpg\uFF0Cjpeg\uFF0Cpng\u683C\u5F0F ")],-1),P=n({name:"userSetup"}),Vt=n({...P,setup(Q){const t=y({defaultAvatar:""}),i=async()=>{try{const o=await h();for(const m in t)t[m]=o[m]}catch(o){console.log("\u83B7\u53D6=>",o)}},l=async()=>{try{await k(t),w.msgSuccess("\u64CD\u4F5C\u6210\u529F"),i()}catch(o){console.log("\u4FDD\u5B58=>",o)}};return i(),(o,m)=>{const c=b,s=B,_=E,d=g,f=A,F=D,C=V("perms");return p(),x("div",j,[e(d,{shadow:"never",class:"!border-none"},{default:r(()=>[I,e(_,{ref:"formRef",model:a(t),"label-width":"120px"},{default:r(()=>[e(s,{label:"\u7528\u6237\u9ED8\u8BA4\u5934\u50CF"},{default:r(()=>[u("div",null,[e(c,{modelValue:a(t).defaultAvatar,"onUpdate:modelValue":m[0]||(m[0]=v=>a(t).defaultAvatar=v),limit:1},null,8,["modelValue"])])]),_:1}),e(s,null,{default:r(()=>[O]),_:1})]),_:1},8,["model"])]),_:1}),S((p(),U(F,null,{default:r(()=>[e(f,{type:"primary",onClick:l},{default:r(()=>[N("\u4FDD\u5B58")]),_:1})]),_:1})),[[C,["setting:user:save"]]])])}}});export{Vt as default};
