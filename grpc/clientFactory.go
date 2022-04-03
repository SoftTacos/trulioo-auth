package grpc

func DoAThing(){

}

//func CreateClient(connections *[]*grpc.ClientConn, opts ...ClientOptions) *grpc.ClientConn {
//	connection, err := createRpcConnection(url, opts...)
//	if err != nil {
//		log.Panic(err)
//	}
//	*connections = append(*connections, connection)
//	return connection
//}
//
//func createRpcConnection(target string, opts ...ClientOptions) (connection *grpc.ClientConn, err error) {
//	logger := logrus.WithFields(logrus.Fields{"clientName": clientName, "target": target})
//	logger.Debug("Creating RPC Client for " + clientName)
//
//	opt := ClientOptions{
//		MaxGrpcServerRecvSize: maxGrpcServerRecvSize,
//		MaxGrpcServerSendSize: maxGrpcServerSendSize,
//		WithPerRetryTimeout:   withPerRetryTimeout,
//	}
//	if len(opts) > 0 {
//		if opts[0].MaxGrpcServerRecvSize > 0 {
//			opt.MaxGrpcServerRecvSize = opts[0].MaxGrpcServerRecvSize
//		}
//		if opts[0].MaxGrpcServerSendSize > 0 {
//			opt.MaxGrpcServerSendSize = opts[0].MaxGrpcServerSendSize
//		}
//		if opts[0].WithPerRetryTimeout > 0 {
//			opt.WithPerRetryTimeout = opts[0].WithPerRetryTimeout
//		}
//	}
//
//	if target == "" || !strings.Contains(target, ":") {
//		logger.Panic(fmt.Sprintf("invalid url given for %s gRPC client: %s", clientName, target))
//	}
//
//	retryOptions := []grpc_retry.CallOption{
//		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
//		grpc_retry.WithCodes(codes.Aborted, codes.Canceled, codes.Unavailable),
//		grpc_retry.WithPerRetryTimeout(opt.WithPerRetryTimeout),
//	}
//
//	connection, err = grpc.Dial(
//		target,
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//		grpc.WithDefaultCallOptions(
//			grpc.MaxCallRecvMsgSize(opt.MaxGrpcServerRecvSize),
//			grpc.MaxCallSendMsgSize(opt.MaxGrpcServerSendSize)),
//		grpc.WithChainUnaryInterceptor(
//			printMethodOnFailedCall,
//			nrgrpc.UnaryClientInterceptor,
//			grpc_retry.UnaryClientInterceptor(retryOptions...)),
//		/*grpc.WithChainStreamInterceptor(
//		)*/
//	)
//
//	if err != nil {
//		logger.WithError(err).Error("could not connect to gRPC client")
//	}
//	return
//}