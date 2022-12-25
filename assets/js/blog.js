let myProject = []

function getData(event){
    event.preventDefault()
    let projectName = document.getElementById("input-project-name").value
    let startDate = document.getElementById("input-start-date").value
    let endDate = document.getElementById("input-end-date").value
    let blogContent = document.getElementById("input-blog-contents").value
    let nodeJsEl = document.getElementById("nodejs")
    let reactJsEl = document.getElementById("reactjs")
    let nextJsEl = document.getElementById("nextjs")
    let typeScriptEl = document.getElementById("typescript")
    let image = document.getElementById("input-blog-image").files

    let nodeJs = (nodeJsEl.checked === true) ? "<i class='fab fa-node-js'></i>" : ""
    let reactJs = (reactJsEl.checked === true) ? "<i class='fab fa-react'></i>" : ""
    let nextJs = (nextJsEl.checked === true) ? "assets/images/nextjs.png" : ""
    let typeScript = (typeScriptEl.checked === true) ? "assets/images/typescr.png" : ""


console.log(image)


    

    if (projectName == "" || startDate == "" || endDate == "" || blogContent == ""
        || image == "") {
            return alert("All input must not be empty!")
        }
    image = URL.createObjectURL(image[0]);
    console.log(image)

    let project = {
        projectName,
        startDate,
        endDate,
        blogContent,
        image,
        nodeJs,
        reactJs,
        nextJs,
        typeScript,
        postedAt : new Date(),
    };
    myProject.push(project);
    renderBlog()



    console.table(myProject)
}

function renderBlog() {
    document.getElementById("contents").innerHTML= "";

    for (let i = 0; i < myProject.length; i++){
        document.getElementById("contents").innerHTML += `
        <div class="list-container">
                    <div class="project-list-item">
                        <div class="project-list-thumbnail">
                            <img src="${myProject[i].image}" alt=""/>
                        </div>
                        <div class="project-list-title">
                            <h4>${myProject[i].projectName}</h4>
                            <p>Durasi: ${getTimeDifference(myProject[i].startDate, myProject[i].endDate)}</p>
                        </div>
                        <div class="project-list-content">
                            <p>${myProject[i].blogContent}</p>
                        </div>
                        <div class="project-list-icon">
                                ${myProject[i].nodeJs}         
                                ${myProject[i].reactJs} 
                                <img src="${myProject[i].nextJs}" alt="">
                                <img src="${myProject[i].typeScript}" alt="">
                        </div>
                        <div class="list-input-container">
                            <div>
                                <button class="project-list-input">edit</button>
                            </div>
                            <div>
                                <button  class="project-list-input">delete</button>
                            </div >
                        </div>
                    </div>
                </div>
        `
    }
}

function getTimeDifference(startDate, endDate){
    let start = new Date(startDate)
    let end = new Date(endDate)

    let ms = 1000

    let timeDifference = end - start

    let differenceMonth = Math.floor(timeDifference / (ms * 3600 * 24 * 30))
    let differenceDay = Math.floor(timeDifference / (ms * 3600 * 24))

    if (differenceMonth > 0){
        return `${differenceMonth} Month Ago`
    } else {
        return `${differenceDay} Days Ago`
    }

} 