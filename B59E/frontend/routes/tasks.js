var express = require("express")
var router = express.Router()
var mongojs = require("mongojs")
var db = mongojs(
    "meantask",
    ["tasks"]
)

router.get("/tasks" , (req, res) => {
    db.tasks.find((err, tasks) => {
        if(err) {
            res.send(err)
        }
        res.json(tasks)
    })
})

router.post("/tasks" , (req, res ) => {
    var task = req.body
    console.log(task)
    if(!task.title){
        res.status(400)
        res.json({
            error: "Bad Data"
        })
    }else{
        db.tasks.save(task, (err, task) => {
            if(err){
                res.send(err)
            }
            res.json(task)
        })
    }
})

router.delete("/task/:id" , (req , res) =>  {
    db.tasks.remove({_id: mongojs.ObjectId(req.params.id)}, (err, task) => {
        if(err) {
            res.send(err)
        }
        res.json(task)
    })
})
module.exports = router