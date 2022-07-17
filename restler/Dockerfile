ARG FUNCTION_DIR="/function"
FROM  mcr.microsoft.com/restlerfuzzer/restler:v7.4.0 as build-image
RUN apk add --no-cache \
    python3 \
    python3-dev \
    libstdc++ \
    build-base \
    libtool \ 
    autoconf \ 
    automake \ 
    libexecinfo-dev \ 
    make \
    cmake \ 
    libcurl



ARG FUNCTION_DIR
RUN mkdir -p ${FUNCTION_DIR}
COPY app/* ${FUNCTION_DIR}
RUN python3 -m pip install --upgrade pip
RUN pip3 install \
        --target ${FUNCTION_DIR} \
        awslambdaric

FROM mcr.microsoft.com/restlerfuzzer/restler:v7.4.0
ARG FUNCTION_DIR
WORKDIR ${FUNCTION_DIR}
# Copy in the build image dependencies
COPY --from=build-image ${FUNCTION_DIR} ${FUNCTION_DIR}
COPY ./entry_script.sh /entry_script.sh
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/local/bin/aws-lambda-rie
RUN chmod +x /entry_script.sh
RUN chmod +x /usr/local/bin/aws-lambda-rie
ENTRYPOINT [ "/entry_script.sh" ]
CMD [ "app.handler" ] 