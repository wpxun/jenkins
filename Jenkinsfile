properties([pipelineTriggers([upstream('go/go-test')])])

node("node02") {
    stage("source"){
        echo 'Hello World'
        echo $id
    }
    stage("build"){
        echo 'Hello World'
    }
    stage("test"){
        echo 'Hello World'
    }
    stage("push"){
        echo 'Hello World'
    }
 
}
