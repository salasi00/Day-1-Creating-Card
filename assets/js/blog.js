let myProject = []
let checkboxes  =  []

function getData(event){
    event.preventDefault()
    let projectName = document.getElementById("input-project-name").value
    let startDate = document.getElementById("input-start-date").value
    let endDate = document.getElementById("input-end-date").value
    let blogContent = document.getElementById("input-blog-contents").value
    let techEl = document.getElementsByName("technologies")
    let image = document.getElementById("input-blog-image")

for (let i = 0; i <techEl.length; i++){
    if (techEl[i].checked){
        checkboxes.push(techEl[i].value)
    }
}



    console.log(image)

    if (projectName == "" || startDate == "" || endDate == "" || blogContent == ""
        || image == "") {
            return alert("All input must not be empty!")
        }
    image = URL.createObjectURL(image.files[0]);

    let project = {
        projectName,
        startDate,
        endDate,
        blogContent,
        image,
        checkboxes,
        postedAt : new Date(),
    };
    myProject.push(project);
    renderBlog()
}

function renderBlog() {
    document.getElementById("contents").innerHTML= "";

    for (let i = 0; i < myProject.length; i++){
        document.getElementById("contents").innerHTML +=
        <div class="list-container">
                    <div class="project-list-item">
                        <div class="project-list-thumbnail">
                            <img src="${myProject[i].image}" alt=""/>
                        </div>
                        <div class="project-list-title">
                            <h4>${myProject[i].projectName}</h4>
                            <p>durasi: 3 Bulan</p>
                        </div>
                        <div class="project-list-content">
                            <p>${myProject[i].blogContent}</p>
                        </div>
                        <div>
                            <span>
                                
                            </span>
                            <span>
                                
                            </span>
                            <span>
                                
                            </span>
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
    }
}