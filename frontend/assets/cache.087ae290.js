import{E as p,I as h}from"./element-plus.078d4249.js";import{s as b}from"./system.ca22106c.js";import{C as r}from"./vue-echarts.ab8d3915.js";import{d as u,r as v,a1 as F,o as f,c as C,X as c,P as d,a as e,V as o,u as t,b7 as E,b6 as A}from"./@vue.a137a740.js";import{d as y}from"./index.1eff73d9.js";import"./@vueuse.07613b64.js";import"./@element-plus.3660753f.js";import"./lodash-es.a31ceab4.js";import"./dayjs.bd523028.js";import"./axios.d8168cfd.js";import"./async-validator.fb49d0f5.js";import"./@ctrl.fd318bfa.js";import"./@popperjs.36402333.js";import"./escape-html.e5dfadb9.js";import"./normalize-wheel-es.8aeb3683.js";import"./resize-detector.4e96b72b.js";import"./echarts.7e912674.js";import"./zrender.754e8e90.js";import"./tslib.60310f1a.js";import"./lodash.48927ea5.js";import"./vue-router.9605b890.js";import"./pinia.9b4180ce.js";import"./css-color-function.b4c88e1a.js";import"./color.a9016252.js";import"./clone.73d1916b.js";import"./color-convert.755d189f.js";import"./color-name.e7a4e1d3.js";import"./color-string.e356f5de.js";import"./balanced-match.d2a36341.js";import"./ms.564e106c.js";import"./nprogress.0f0f7ca7.js";import"./vue-clipboard3.4e164ffd.js";import"./clipboard.7c3d630c.js";import"./highlight.js.7165574c.js";import"./@highlightjs.7fc78ec7.js";const l=i=>(E("data-v-4f669e5f"),i=i(),A(),i),D={class:"cache"},B=l(()=>e("div",{class:"mb-4 lg"},"\u57FA\u672C\u4FE1\u606F",-1)),w={class:"el-table--enable-row-transition el-table--large el-table"},x={class:"el-table__body",cellspacing:"0"},g={class:"el-table__row"},S=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"Redis\u7248\u672C")],-1)),O={class:"el-table__cell"},I={class:"cell"},k=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"\u8FD0\u884C\u6A21\u5F0F")],-1)),z={class:"el-table__cell"},P={class:"cell"},V=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"\u7AEF\u53E3")],-1)),N={class:"el-table__cell"},R={class:"cell"},X=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"\u5BA2\u6237\u7AEF\u6570")],-1)),K={class:"el-table__cell"},L={class:"cell"},M={class:"el-table__row"},U=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"\u8FD0\u884C\u65F6\u95F4(\u5929)")],-1)),j={class:"el-table__cell"},q={class:"cell"},G=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"\u4F7F\u7528\u5185\u5B58")],-1)),H={class:"el-table__cell"},J={class:"cell"},Q=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"\u4F7F\u7528CPU")],-1)),T={class:"el-table__cell"},W={class:"cell"},Y=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"\u5185\u5B58\u914D\u7F6E")],-1)),Z={class:"el-table__cell"},$={class:"cell"},ee={class:"el-table__row"},se=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"AOF\u662F\u5426\u5F00\u542F")],-1)),te={class:"el-table__cell"},le={class:"cell"},oe=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"RDB\u662F\u5426\u6210\u529F")],-1)),ce={class:"el-table__cell"},ae={class:"cell"},_e=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"Key\u6570\u91CF")],-1)),ie={class:"el-table__cell"},de={class:"cell"},ne=l(()=>e("td",{class:"el-table__cell"},[e("div",{class:"cell"},"\u7F51\u7EDC\u5165\u53E3/\u51FA\u53E3")],-1)),re={class:"el-table__cell"},ue={class:"cell"},me={class:"sm:flex"},pe=l(()=>e("div",{class:"mb-10"},"\u547D\u4EE4\u7EDF\u8BA1",-1)),he={class:"flex h-[300px] items-center"},be=l(()=>e("div",{class:"mb-10"},"\u5185\u5B58\u4FE1\u606F",-1)),ve={class:"flex h-[300px] items-center"},Fe=u({name:"cache"}),fe=u({...Fe,setup(i){const s=v({}),a=F({commandChartOption:{tooltip:{trigger:"item"},series:[{label:{show:!0},labelLine:{show:!0},type:"pie",radius:"85%",color:["#0D47A1","#1565C0","#1976D2","#1E88E5","#2196F3","#42A5F5","#64B5F6","#90CAF9","#BBDEFB","#E3F2FD","#CAF0F8","#ADE8F4","#90E0EF","#48CAE4","#00B4D8","#0096C7","#0077B6","#023E8A","#03045E","#8ecae6","#98c1d9","#D9ED92","#B5E48C","#99D98C","#76C893","#52B69A","#34A0A4","#168AAD","#1A759F","#1E6091","#184E77","#457b9d"],data:[{value:"",name:""}],emphasis:{itemStyle:{shadowBlur:10,shadowOffsetX:0,shadowColor:"rgba(0, 0, 0, 0.5)"}}}]},memoryChartOption:{tooltip:{formatter:"{a} <br/>{b} : {c}%"},series:[{name:"Pressure",type:"gauge",radius:"100%",detail:{formatter:"{value}"},data:[{value:"",name:"\u5185\u5B58\u6D88\u8017"}]}]}});return(async()=>{const _=await b();s.value=_.info,s.value.dbSize=_.dbSize,a.commandChartOption.series[0].data=_.commandStats,a.memoryChartOption.series[0].data[0].value=(_.info.used_memory/1024/1024).toFixed(2),a.memoryChartOption.series[0].detail.formatter="{value}M"})(),(_,Ee)=>{const m=p,n=h;return f(),C("div",D,[c(n,{class:"!border-none",shadow:"never"},{default:d(()=>[e("div",null,[B,e("div",w,[c(m,null,{default:d(()=>[e("table",x,[e("tbody",null,[e("tr",g,[S,e("td",O,[e("div",I,o(t(s).redis_version),1)]),k,e("td",z,[e("div",P,o(t(s).redis_mode=="standalone"?"\u5355\u673A":"\u96C6\u7FA4"),1)]),V,e("td",N,[e("div",R,o(t(s).tcp_port),1)]),X,e("td",K,[e("div",L,o(t(s).connected_clients),1)])]),e("tr",M,[U,e("td",j,[e("div",q,o(t(s).uptime_in_days),1)]),G,e("td",H,[e("div",J,o(t(s).used_memory_human),1)]),Q,e("td",T,[e("div",W,o(t(s).used_cpu_user_children),1)]),Y,e("td",Z,[e("div",$,o(t(s).maxmemory_human),1)])]),e("tr",ee,[se,e("td",te,[e("div",le,o(t(s).aof_enabled==0?"\u5F00\u542F":"\u5173\u95ED"),1)]),oe,e("td",ce,[e("div",ae,o(t(s).aof_enabled=="ok"?"\u6210\u529F":"\u5931\u8D25"),1)]),_e,e("td",ie,[e("div",de,o(t(s).dbSize),1)]),ne,e("td",re,[e("div",ue,o(t(s).instantaneous_input_kbps)+" / "+o(t(s).instantaneous_output_kbps),1)])])])])]),_:1})])])]),_:1}),e("div",me,[c(n,{class:"sm:mr-4 flex-1 !border-none mt-4",shadow:"never"},{default:d(()=>[e("div",null,[pe,e("div",he,[c(t(r),{autoresize:"",option:a.commandChartOption},null,8,["option"])])])]),_:1}),c(n,{class:"flex-1 !border-none mt-4",shadow:"never"},{default:d(()=>[e("div",null,[be,e("div",ve,[c(t(r),{autoresize:"",option:a.memoryChartOption},null,8,["option"])])])]),_:1})])])}}});const ls=y(fe,[["__scopeId","data-v-4f669e5f"]]);export{ls as default};
