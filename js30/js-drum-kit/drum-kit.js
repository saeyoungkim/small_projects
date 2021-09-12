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

const instruments = document.getElementsByClassName('instrument')

const twinkleAndPlayOnce = (ch) => {
    const i = keyIndexMap[ch];

    const audio = new Audio(`./audios/${instruments[i].dataset['key']}.wav`)

    instruments[i].classList.add('clicked')
    audio.play();
    setTimeout(() => {
        instruments[i].classList.remove('clicked');
        audio.remove();
    }, 50);
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