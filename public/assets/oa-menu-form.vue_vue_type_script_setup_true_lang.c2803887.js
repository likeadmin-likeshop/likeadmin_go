import{O as C,P as g,C as B,D as I,F as P}from"./element-plus.078d4249.js";import{r as k}from"./useMenuOa.e15574fe.js";import{d as w,s as U,r as c,w as R,l as x,o as m,O as F,P as l,X as u,u as a,U as f,T as i,K as A,c as v,W as E}from"./@vue.a137a740.js";const G=w({__name:"oa-menu-form",props:{modular:{default:"master"},name:{default:""},menuType:{default:1},visitType:{default:"view"},url:{default:""},appId:{default:""},pagePath:{default:""}},emits:["update:name","update:menuType","update:visitType","update:url","update:pagePath"],setup(r,{expose:b,emit:d}){const y=r,_=U(),e=c({...y});return R(()=>y,V=>{e.value=V},{immediate:!0}),x(()=>{y.modular==="master"&&d("update:menuType",e.value.menuType),d("update:name",e.value.name),d("update:visitType",e.value.visitType),d("update:url",e.value.url),d("update:appId",e.value.appId),d("update:pagePath",e.value.pagePath)}),b({menuFormRef:_}),(V,t)=>{const n=B,p=I,s=C,T=g,D=P;return m(),F(D,{ref_key:"menuFormRef",ref:_,rules:a(k),model:a(e),"label-width":"120px",class:"pr-10"},{default:l(()=>[u(p,{label:r.modular==="master"?"\u4E3B\u83DC\u5355\u540D\u79F0":"\u5B50\u83DC\u5355\u540D\u79F0",prop:"name"},{default:l(()=>[u(n,{modelValue:a(e).name,"onUpdate:modelValue":t[0]||(t[0]=o=>a(e).name=o)},null,8,["modelValue"])]),_:1},8,["label"]),r.modular==="master"?(m(),F(p,{key:0,label:"\u4E3B\u83DC\u5355\u7C7B\u578B",prop:"menuType"},{default:l(()=>[u(T,{modelValue:a(e).menuType,"onUpdate:modelValue":t[1]||(t[1]=o=>a(e).menuType=o)},{default:l(()=>[u(s,{label:1},{default:l(()=>[f("\u4E0D\u914D\u7F6E\u5B50\u83DC\u5355")]),_:1}),u(s,{label:2},{default:l(()=>[f("\u914D\u7F6E\u5B50\u83DC\u5355")]),_:1})]),_:1},8,["modelValue"])]),_:1})):i("",!0),a(e).menuType===2&&r.modular==="master"?(m(),F(p,{key:1,label:""},{default:l(()=>[A(V.$slots,"default")]),_:3})):i("",!0),a(e).menuType===1?(m(),v(E,{key:2},[u(p,{label:"\u8DF3\u8F6C\u94FE\u63A5",prop:"visitType"},{default:l(()=>[u(T,{modelValue:a(e).visitType,"onUpdate:modelValue":t[2]||(t[2]=o=>a(e).visitType=o)},{default:l(()=>[u(s,{label:"view"},{default:l(()=>[f("\u7F51\u9875")]),_:1}),u(s,{label:"miniprogram"},{default:l(()=>[f("\u5C0F\u7A0B\u5E8F")]),_:1})]),_:1},8,["modelValue"])]),_:1}),u(p,{label:"\u7F51\u5740",prop:"url"},{default:l(()=>[u(n,{modelValue:a(e).url,"onUpdate:modelValue":t[3]||(t[3]=o=>a(e).url=o)},null,8,["modelValue"])]),_:1}),a(e).visitType=="miniprogram"?(m(),v(E,{key:0},[u(p,{label:"AppId",prop:"appId"},{default:l(()=>[u(n,{modelValue:a(e).appId,"onUpdate:modelValue":t[4]||(t[4]=o=>a(e).appId=o)},null,8,["modelValue"])]),_:1}),u(p,{label:"\u8DEF\u5F84",prop:"pagePath"},{default:l(()=>[u(n,{modelValue:a(e).pagePath,"onUpdate:modelValue":t[5]||(t[5]=o=>a(e).pagePath=o)},null,8,["modelValue"])]),_:1})],64)):i("",!0)],64)):i("",!0)]),_:3},8,["rules","model"])}}});export{G as _};
