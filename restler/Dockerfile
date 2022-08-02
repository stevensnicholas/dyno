ARG FUNCTION_DIR="/function"
FROM  mcr.microsoft.com/restlerfuzzer/restler:v8.5.0 as build-image
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
    awslambdaric \
    boto3
FROM mcr.microsoft.com/restlerfuzzer/restler:v8.5.0
ARG FUNCTION_DIR
WORKDIR ${FUNCTION_DIR}
COPY --from=build-image ${FUNCTION_DIR} ${FUNCTION_DIR}
COPY ./entry_script.sh /entry_script.sh
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/download/v1.6/aws-lambda-rie /usr/local/bin/aws-lambda-rie
RUN chmod +x /usr/local/bin/aws-lambda-rie
ENTRYPOINT [ "/entry_script.sh" ]
CMD [ "app.handler" ] 