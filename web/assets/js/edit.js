window.onload = function () {
    var contenteditableDiv = document.getElementById("textdiv");
    
    contenteditableDiv.addEventListener("input", function(){ 
        document.getElementById("text").value = contenteditableDiv.innerText;
    });

    contenteditableDiv.focus();
};