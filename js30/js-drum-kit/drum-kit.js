const keyIndexMap = {
    'a': 0,
    's': 1,
    'd': 2,
    'f': 3,
    'g': 4,
    'h': 5,
    'j': 6,
    'k': 7
};

const instruments = document.querySelectorAll('.instrument');
const audios = Array.from(instruments).map(elem => new Audio(`./audios/${elem.dataset['key']}.wav`));

const removeTransition = e => {
    if(e.propertyName !== 'transform') {
        return ;
    }

    // this means the caller of event listener
    e.srcElement.classList.remove('playing');
}

instruments.forEach(elem => elem.addEventListener('transitionend', removeTransition));

const twinkleAndPlayOnce = (ch) => {
    const i = keyIndexMap[ch];

    instruments[i].classList.add('playing');
    audios[i].currentTime = 0;
    audios[i].play();
}

document.addEventListener('keypress', (e) => {
    if(e.key === undefined) {
        return ;
    }

    switch (e.key) {
        case 'a':
        case 'A':
            twinkleAndPlayOnce('a');
            break;
        case 's':
        case 'S':
            twinkleAndPlayOnce('s');
            break;
        case 'd':
        case 'D':
            twinkleAndPlayOnce('d');
            break;
        case 'f':
        case 'F':
            twinkleAndPlayOnce('f');
            break;
        case 'g':
        case 'G':
            twinkleAndPlayOnce('g');
            break;
        case 'h':
        case 'H':
            twinkleAndPlayOnce('h');
            break;
        case 'j':
        case 'J':
            twinkleAndPlayOnce('j');
            break;
        case 'k':
        case 'K':
            twinkleAndPlayOnce('k');
            break;
        default:
            return;
    } 
});