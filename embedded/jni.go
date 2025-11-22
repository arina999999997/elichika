package embedded

// #cgo LDFLAGS: -llog
// #include <android/log.h>
// #include <stdlib.h>
// #include <jni.h>
// #include <unistd.h>
//
// JavaVM* jvm;
// jclass unityPlayerClass;
// jmethodID elichikaAlertMethod;
// JNIEnv* env;
//
// void show_android_alert(char* message) {
//    // show alert
//    jint getEnvResult = (*jvm)->GetEnv(jvm, (void **)&env, JNI_VERSION_1_6);
//    if (getEnvResult == JNI_EDETACHED) {
//        if ((*jvm)->AttachCurrentThread(jvm, &env, NULL) != 0) {
//            __android_log_print(ANDROID_LOG_FATAL, "elichika", "failed to attach the env");
//        }
//    }
//    jstring jstr = (*env)->NewStringUTF(env, message);
//    (*env)->CallStaticVoidMethod(env, unityPlayerClass, elichikaAlertMethod, jstr);
// }
//
// JNIEXPORT jint JNI_OnLoad(JavaVM* vm, void* reserved) {
//     jvm = vm;
//     jint getEnvResult = (*jvm)->GetEnv(jvm, (void **)&env, JNI_VERSION_1_6);
//     if (getEnvResult == JNI_EDETACHED) {
//         if ((*jvm)->AttachCurrentThread(jvm, &env, NULL) != 0) {
//             __android_log_print(ANDROID_LOG_FATAL, "elichika", "failed to attach the env");
//         }
//     }
//     unityPlayerClass = (*env)->FindClass(env, "com/unity3d/player/UnityPlayer");
//     elichikaAlertMethod = (*env)->GetStaticMethodID(env, unityPlayerClass, "elichikaAlert", "(Ljava/lang/String;)V");
//     if (elichikaAlertMethod == NULL) {
//         __android_log_print(ANDROID_LOG_FATAL, "elichika", "failed to find elichikaAlert method, alert will not be shown");
//     }
//     return JNI_VERSION_1_6;
// }
//
// JNIEXPORT void JNI_OnUnload(JavaVM* vm, void* reserved) {
//
// }
import "C"

import "unsafe"

func Alert(s string) {
	cString := C.CString(s)
	C.show_android_alert(cString)
	C.free(unsafe.Pointer(cString))
}

func SendLoadedSignal() {
	Alert("Elichika loading finished!")
}
