(getInterfaceList = async () => {
  "use strict";

  let url      = "http://localhost:8989/interfaces";
  let response = await fetch(url, {
    "method": "GET"
  });
  let json    = await response.json();
  let element = document.getElementById("divId01");
  
  if (element.hasChildNodes() == true) {

    let childNodes = element.childNodes;
    
    for (let i=0; i<childNodes.length; i++) {
      if (childNodes[i].id == "interfaceList") {
        element.removeChild(childNodes[i])
      }
    }
  }

  let ol = document.createElement("ol");
  ol.setAttribute("id", "interfaceList");

  
  json.forEach(json_element => {
    let li   = document.createElement("li");
    let html = "";

    for (let key in json_element) {
      html += key + " : " + json_element[key] + "<br>";
    }

    li.innerHTML = html

    ol.appendChild(li)
  });

  element.appendChild(ol)
})();