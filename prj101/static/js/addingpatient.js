function addPatient(){
    console.log("dhendup")
    var _data = {
        patientname : document.getElementById("patientname").value,
        cid : document.getElementById("cid").value,
        age : document.getElementById("age").value,
        address : document.getElementById("address").value,
        disease : document.getElementById("disease").value,
        date : document.getElementById("date").value

    }
    // console.log(document.getElementById("patientname").value)
    // alert("oiiiiii")

   
    fetch('/patients/add',{
        method: "POST",
        body: JSON.stringify(_data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    })  
}
function loggingout(){
    if(confirm("are tou sure")){
        location.href="index.html"
    }
}

// patient updating
var __data = null;
function SearchingPatient(){
    document.getElementById("add").classList.add("d-none")
    document.getElementById("update").classList.remove("d-none")
    document.getElementById("delete").classList.remove("d-none")
    __data = {
        cid : document.getElementById("search").value,
    }
    // console.log("dhendup")
    // console.log(__data.cid)
    var cid = __data.cid
    fetch('/patients/get/'+cid)
    .then(res => res.text())
    .then(data => b(data))
}
function b(credentials){
    const data = JSON.parse(credentials)
    if(__data.cid != data.CID){
        alert("there is no such person")
        document.getElementById("search").value =""
        return
    }
    document.getElementById("patientname").value = data.PatientName
    document.getElementById("cid").value=data.CID
    document.getElementById("age").value = data.Age
    document.getElementById("address").value =data.Address
    document.getElementById("disease").value = data.Disease
    document.getElementById("date").value = data.Date
    document.getElementById("search").value =""
    __data = null
}
function UpdatingPatient(){
    __data ={
        patientname : document.getElementById("patientname").value,
        cid : document.getElementById("cid").value,
        age : document.getElementById("age").value ,
        address : document.getElementById("address").value,
        disease : document.getElementById("disease").value,
        date : document.getElementById("date").value 
    }
    console.log(__data.patientname)
    console.log(__data.cid)
    console.log(__data.age)
    console.log(__data.address)
    console.log(__data.disease)
    console.log(__data.date)
    var cid = __data.cid

    fetch('/patients/update/'+cid,{
        method: "PUT",
        body: JSON.stringify(__data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    })
    __data = null
}


function Deleting(){
    if(confirm("Are you sure ?")){
        __data ={
            cid : document.getElementById("cid").value
        }
        fetch('patients/delete/'+__data.cid,{
            method : "DELETE",
            headers: {"Content-type": "application/json; charset=UTF-8"}
        }) 
    }
}

