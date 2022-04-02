var translation = {
    "0": {
        "name": "Zero",
        "phon": "(ZEE-RO)"
    },
    "1": {
        "name": "One",
        "phon": "(WUN)"
    },
    "2": {
        "name": "Two",
        "phon": "(TOO)"
    },
    "3": {
        "name": "Three",
        "phon": "(TREE)"
    },
    "4": {
        "name": "Four",
        "phon": "(FOW-ER)"
    },
    "5": {
        "name": "Five",
        "phon": "(FIFE)"
    },
    "6": {
        "name": "Six",
        "phon": "(SIX)"
    },
    "7": {
        "name": "Seven",
        "phon": "(SEV-EN)"
    },
    "8": {
        "name": "Eight",
        "phon": "(AIT)"
    },
    "9": {
        "name": "Nine",
        "phon": "(NIN-ER)"
    },
    "A": {
        "name": "Alfa",
        "phon": "(AL-FAH)"
    },
    "B": {
        "name": "Bravo",
        "phon": "(BRAH-VOH)"
    },
    "C": {
        "name": "Charlie",
        "phon": "(CHAR-LEE) or (SHAR-LEE)"
    },
    "D": {
        "name": "Delta",
        "phon": "(DELL-TAH)"
    },
    "E": {
        "name": "Echo",
        "phon": "(ECK-OH)"
    },
    "F": {
        "name": "Foxtrot",
        "phon": "(FOKS-TROT)"
    },
    "G": {
        "name": "Golf",
        "phon": "(GOLF)"
    },
    "H": {
        "name": "Hotel",
        "phon": "(HOH-TEL)"
    },
    "I": {
        "name": "India",
        "phon": "(IN-DEE-AH)"
    },
    "J": {
        "name": "Juliett",
        "phon": "(JEW-LEE-ETT)"
    },
    "K": {
        "name": "Kilo",
        "phon": "(KEY-LOH)"
    },
    "L": {
        "name": "Lima",
        "phon": "(LEE-MAH)"
    },
    "M": {
        "name": "Mike",
        "phon": "(MIKE)"
    },
    "N": {
        "name": "November",
        "phon": "(NO-VEM-BER)"
    },
    "O": {
        "name": "Oscar",
        "phon": "(OSS-CAH)"
    },
    "P": {
        "name": "Papa",
        "phon": "(PAH-PAH)"
    },
    "Q": {
        "name": "Quebec",
        "phon": "(KEH-BECK)"
    },
    "R": {
        "name": "Romeo",
        "phon": "(ROW-ME-OH)"
    },
    "S": {
        "name": "Sierra",
        "phon": "(SEE-AIR-RAH)"
    },
    "T": {
        "name": "Tango",
        "phon": "(TANG-GO)"
    },
    "U": {
        "name": "Uniform",
        "phon": "(YOU-NEE-FORM) or (OO-NEE-FORM)"
    },
    "V": {
        "name": "Victor",
        "phon": "(VIK-TAH)"
    },
    "W": {
        "name": "Whiskey",
        "phon": "(WISS-KEY)"
    },
    "X": {
        "name": "Xray",
        "phon": "(ECKS-RAY)"
    },
    "Y": {
        "name": "Yankee",
        "phon": "(YANG-KEY)"
    },
    "Z": {
        "name": "Zulu",
        "phon": "(ZOO-LOO)"
    }
};

var app = Vue.createApp({
  data: function () {
    return {
      message: '',
      phonetic: 'No message'
    }
  },
  methods: {
    updateMessage: function () {
      var msg = this.message.toUpperCase();

      var output = [];
      for (var i = 0; i < msg.length; i++) {
        var c = msg.charAt(i);
        if (translation[c] !== undefined) {
          // TODO: Do something with the pronunciations of each letter.
          output.push(translation[c].name);
        } else {
          output.push(c);
        }
      }

      this.phonetic = output.join(' ');
    }
  }
});

const vm = app.mount("#app");
