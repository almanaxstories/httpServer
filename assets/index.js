console.log('JS-file connected successfully!');

window.addEventListener('load', ()=>{
let appContainer = document.getElementById('signupFormDiv');
let signupForm = newElement('FORM');
signupForm.method = 'POST';
signupForm.setAttribute('class', 'form');
appContainer.append(signupForm);
let firstNameLabel = newElement("label", "", "First Name:", "fNameText", "fNameLabel");
signupForm.append(firstNameLabel);
let firstNameInput = newElement("input", "text", "", "", "fNameText");
firstNameInput.setAttribute("name","fNameText");
signupForm.append(firstNameInput);
let lastNameLabel = newElement("label", "", "Last Name:", "lNameText", "lNameLabel");
signupForm.append(lastNameLabel);
let lastNameInput = newElement("input", "text", "", "", "lNameText");
lastNameInput.setAttribute("name","lNameText");
signupForm.append(lastNameInput);
let submitButton = newElement("input", "button", "Let's go!", "", "submitButton");
signupForm.append(submitButton);

submitButton.addEventListener('click', event => {
    event.preventDefault();
    const formData = new FormData(signupForm);
    const objectFromInputs = Object.fromEntries(formData);

    fetch('http://0.0.0.0:8090/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        },
        body: JSON.stringify(objectFromInputs)
    }).then(res => res.json()).then(data => console.log(data)).catch(err => console.log(err));
    //

});
});

function newElement(tag, typeOfInput, text, labelFor, id){
    let element = document.createElement(tag);

    if (typeOfInput){
        element.type = typeOfInput;
    }
    
    if (text){
        element.innerText = text;
    }
    
    if (labelFor){
        element.for = labelFor;
    }

    if (id){
        element.id = id;
    }    
    return element;
}

