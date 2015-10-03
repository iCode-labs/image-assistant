chrome.contextMenus.create({
  title: "收藏图片",
  contexts: ["image"],
  onclick: function(info) {

    handleImageURL(info.srcUrl);
  }

});

function handleImageURL(url) {
  var name = prompt("输入图片的自定义名称");
  var xhr = new XMLHttpRequest();
  var reqUrl = "http://www.mirana.me:8080/store?url=" + url;
  if (name)
    reqUrl += "&name=" + name;
  xhr.open("get", reqUrl);
  xhr.setRequestHeader("content-type", "x-www-form-urlencoded");
  xhr.send();
}
