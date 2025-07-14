import express from 'express';
import { engine } from 'express-handlebars';
import axios from 'axios';

const app = express();

app.engine('handlebars',engine());
app.set("view engine","handlebars");
app.set("views","./views");

// Arquivos CSS e JS
app.use(express.static('./public'))

app.get("/",async (req,res)=>{
    let response = await axios.get('http://localhost:3000/');
    console.log(response.data.msg);
    res.render("home",response.data);
});

app.listen(4000);
