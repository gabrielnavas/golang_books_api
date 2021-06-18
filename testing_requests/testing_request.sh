make_request() {
    URL=$1
    NAME_PARAM=$2
    curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"name":"'$2'"}' \
    $URL
    # http://localhost:8000/category
}

send_request() {
    URL=$1
    END=$2
    for ((i=0; i < END; i++)); do
        param="i am a string number $i"
        echo "$param"
        make_request $URL $param &
    done
}

send_request $1 $2