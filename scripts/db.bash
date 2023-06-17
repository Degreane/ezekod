#!/bin/bash
## DBCheck 
## running environment : docker mongodb

DOCKER=$(which docker)
resp=$?
if [[ ${resp} -eq 0 ]]; then
    ## Check if ezekod docker is initiated
    ${DOCKER} inspect ezekod 1>&2 >/dev/null
    resp=$?
    if [[ ${resp} -eq 0 ]]; then
        ## if it is initiated then we check for the status of the docker file
        JQ=$(which jq)
        resp=$?
        if [[ ${resp} -eq 0 ]]; then 
            Running=$(${DOCKER} inspect ezekod | ${JQ} '.[]|.State["Running"]')
            if [[ ${Running} == true ]]; then 
                ## if it is RUnning then just echo Running 
                printf '%s' "EzeKod Running "
            else 
                ## start docker server
                ${DOCKER} start ezekod
            fi
        else
            printf '%s\n' "'jq' not found please install" 
        fi
    else
        ## Here we should Initiate docker 
        printf '%s' "Starting EzeKod MongoDB Docker"
        ${DOCKER} run -p 127.0.0.1:27017:27017 --name ezekod --detach mongo:latest 
        #db.createUser({"user":"psycho","pwd":"shta2telik",roles:[{role:"readWriteAnyDatabase",db:"admin"},{role:"dbAdminAnyDatabase",db:"admin"}],mechanisms: [ 'SCRAM-SHA-1', 'SCRAM-SHA-256' ]})
        #security:
        #    authorization: enabled
    fi
else
    printf '%s\n%s' "Docker binaries not found" "Please Install docker"
fi

