const h = document.getElementsByClassName('hour')[0];
const m = document.getElementsByClassName('minute')[0];
const s = document.getElementsByClassName('second')[0];

const updateTime = () => {
    const date = new Date();
    
    const secDeg = date.getSeconds() * 6; 
    s.style.transform = `translate(-50%, -100%) rotate(${secDeg}deg)`;
    const minDeg = date.getMinutes() * 6 + secDeg / 60.0;
    m.style.transform = `translate(-50%, -100%) rotate(${minDeg}deg)`;
    const hourDeg = date.getHours() * 30 + secDeg / 720.0;
    h.style.transform = `translate(-50%, -100%) rotate(${hourDeg}deg)`;
}

setInterval(updateTime,250);
// init
updateTime();