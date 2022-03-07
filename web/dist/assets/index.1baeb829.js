import{i as d,r as v,o as u,c as l,w as g,v as h,a,u as w,p as b,b as y,W as p,C as I,B as f,d as $,e as C,f as x,g as k}from"./vendor.ef7b5f8c.js";const D=function(){const o=document.createElement("link").relList;if(o&&o.supports&&o.supports("modulepreload"))return;for(const e of document.querySelectorAll('link[rel="modulepreload"]'))i(e);new MutationObserver(e=>{for(const t of e)if(t.type==="childList")for(const r of t.addedNodes)r.tagName==="LINK"&&r.rel==="modulepreload"&&i(r)}).observe(document,{childList:!0,subtree:!0});function s(e){const t={};return e.integrity&&(t.integrity=e.integrity),e.referrerpolicy&&(t.referrerPolicy=e.referrerpolicy),e.crossorigin==="use-credentials"?t.credentials="include":e.crossorigin==="anonymous"?t.credentials="omit":t.credentials="same-origin",t}function i(e){if(e.ep)return;e.ep=!0;const t=s(e);fetch(e.href,t)}};D();var m=(n,o)=>{for(const[s,i]of o)n[s]=i;return n};const _=n=>(b("data-v-4dc56c61"),n=n(),y(),n),N={class:"form-signin"},S=_(()=>a("h1",{class:"h3 mb-3 font-weight-normal"},"Pin IPFS CID",-1)),A=_(()=>a("label",{for:"inputCid",class:"sr-only"},"CID",-1)),P={setup(n){d("$address");const o=v(""),s=d("$cfg"),i=["function transfer(address to, uint amount) returns (bool)","function allowance(address owner, address spender) view returns (uint256)","function approve(address spender, uint256 amount) returns (bool)"];var e=function(){var t=new p(window.ethereum);new I(s["here.token"],i,t.getSigner()).transfer(s["here.bridge"],f.from(o.value).mul(f.from(10).pow(18)))};return(t,r)=>(u(),l("form",N,[S,A,g(a("input",{"onUpdate:modelValue":r[0]||(r[0]=c=>o.value=c),type:"text",id:"inputCid",class:"mb-3 form-control",placeholder:"The CID of the IPFS file to be pinned",required:"",autofocus:""},null,512),[[h,o.value]]),a("button",{type:"button",class:"btn btn-lg btn-warning btn-block",onClick:r[1]||(r[1]=c=>w(e)())},"PIN ")]))}};var B=m(P,[["__scopeId","data-v-4dc56c61"]]);const L={components:{Deposit:B}},q={role:"main",class:"container",id:"box"};function O(n,o,s,i,e,t){const r=$("Deposit");return u(),l("div",null,[a("main",q,[C(r)])])}var F=m(L,[["render",O]]);x.get("/api/v0/pin/config").then(async function(n){let o=n.data;await window.ethereum.request({method:"eth_requestAccounts"});const s=new p(window.ethereum);await s.getNetwork(),window.ethereum.on("chainChanged",function(){window.location.reload()});var i=await s.getSigner().getAddress();let e=k(F);e.provide("$address",i),e.provide("$cfg",o),e.mount("#app")});
